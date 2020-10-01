package comment

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/comment"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Comment) Create(ctx context.Context, msg *comment.CreateRequest) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId == 0 {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	customer, err := c.customerCli.Get(ctx, &empty.Empty{})
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			return nil, server.NewErrServer(grpcStatus.Code(), errors.WithStack(err))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	if _, err := c.gatewayDB.
		Insert(objectComment).
		Columns(fieldProductId, fieldCustomerId, fieldName, fieldText).
		Values(msg.ProductId, session.CustomerId, customer.Name, msg.Text).
		Exec(); err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
