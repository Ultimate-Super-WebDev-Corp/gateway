package customer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

func (Customer) Create(ctx context.Context, msg *customer.CreateRequest) (*empty.Empty, error) {
	return nil, nil
}
