package review

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (r Review) CreateCommentFromSource(ctx context.Context, msg *review.CreateCommentFromSourceRequest) (*empty.Empty, error) {
	if msg.Source == sourceAggregated || msg.Source == sourceCustomer {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("must not be customer or aggregated"))
	}

	session := server.SessionFromCtx(ctx)
	if !server.IsSessionRoot(session) {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("permission denied"))
	}

	msg.Source = strings.ToLower(msg.Source)
	msg.Source = strings.Join(strings.Fields(msg.Source), "")
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
		if err = r.updateRating(tx, msg.ProductId, msg.Source, msg.Rating, 1); err != nil {
			return err
		}
		if err = r.updateAggregatedRating(tx, msg.ProductId); err != nil {
			return err
		}
		if _, err := r.statementBuilder.
			Insert(objectComment).
			Columns(fieldProductId, fieldName, fieldText, fieldRating, fieldSource).
			Values(msg.ProductId, msg.Name, msg.Text, msg.Rating, msg.Source).
			RunWith(r.gatewayDB).Exec(); err != nil {
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
