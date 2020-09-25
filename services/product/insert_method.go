package product

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

func (p Product) Insert(_ context.Context, msg *product.ProductMsg) (*empty.Empty, error) {
	_, err := p.gatewayDB.
		Insert(objectProduct).
		Columns(fieldName, fieldBrand, fieldDescription, fieldUpdatedAt).
		Values(msg.Name, msg.Brand, msg.Description, time.Now()).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
