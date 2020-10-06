package review

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (r Review) CreateComment(ctx context.Context, msg *review.CreateCommentRequest) (*empty.Empty, error) {
	if msg.Rating == review.Rating_UNDEFINED {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("rating must not be UNDEFINED"))
	}

	session := server.SessionFromCtx(ctx)
	if !server.IsSessionLoggedIn(session) {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	customer, err := r.customerCli.Get(ctx, &empty.Empty{})
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			return nil, server.NewErrServer(grpcStatus.Code(), errors.WithStack(err))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	tx, err := r.gatewayDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	err = func() (err error) {
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			}
		}()
		if err = r.updateRating(tx, msg.ProductId, sourceCustomer, msg.Rating, 1); err != nil {
			return err
		}
		if err = r.updateAggregatedRating(tx, msg.ProductId); err != nil {
			return err
		}
		if _, err := r.statementBuilder.
			Insert(objectComment).
			Columns(fieldProductId, fieldCustomerId, fieldName, fieldText, fieldRating, fieldSource).
			Values(msg.ProductId, session.CustomerId, customer.Name, msg.Text, msg.Rating, sourceCustomer).
			RunWith(tx).Exec(); err != nil {
			return errors.WithStack(err)
		}
		if err = tx.Commit(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
