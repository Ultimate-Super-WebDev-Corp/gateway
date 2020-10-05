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

func (c Review) CreateComment(ctx context.Context, msg *review.CreateCommentRequest) (*empty.Empty, error) {
	if msg.Rating == review.Rating_UNDEFINED {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("rating must not be UNDEFINED"))
	}

	session := server.SessionFromCtx(ctx)
	if !server.IsSessionLoggedIn(session) {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	customer, err := c.customerCli.Get(ctx, &empty.Empty{})
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			return nil, server.NewErrServer(grpcStatus.Code(), errors.WithStack(err))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	//todo update product rating

	if _, err := c.gatewayDB.
		Insert(objectComment).
		Columns(fieldProductId, fieldCustomerId, fieldName, fieldText, fieldRating, fieldSource).
		Values(msg.ProductId, session.CustomerId, customer.Name, msg.Text, msg.Rating, "").
		Exec(); err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
