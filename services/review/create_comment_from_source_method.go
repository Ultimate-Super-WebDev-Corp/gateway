package review

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Review) CreateCommentFromSource(ctx context.Context, msg *review.CreateCommentFromSourceRequest) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	if !server.IsSessionRoot(session) {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("permission denied"))
	}

	//todo update product rating

	if _, err := c.gatewayDB.
		Insert(objectComment).
		Columns(fieldProductId, fieldName, fieldText, fieldRating, fieldSource).
		Values(msg.ProductId, msg.Name, msg.Text, msg.Rating, msg.Source).
		Exec(); err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
