package product

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) Create(ctx context.Context, msg *product.ProductMsg) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	if !server.IsSessionRoot(session) {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("permission denied"))
	}
	_, err := p.gatewayDB.
		Insert(objectProduct).
		Columns(fieldName, fieldBrand, fieldDescription, fieldImages, fieldCountry, fieldUpdatedAt).
		Values(msg.Name, msg.Brand, msg.Description, pq.Array(msg.Images), msg.Country, time.Now().UTC()).Exec()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
