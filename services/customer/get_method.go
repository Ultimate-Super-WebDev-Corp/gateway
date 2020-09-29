package customer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

func (Customer) Get(context.Context, *empty.Empty) (*customer.CustomerMsg, error) {
	return nil, nil
}
