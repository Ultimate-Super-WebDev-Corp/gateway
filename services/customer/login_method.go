package customer

import (
	"context"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

func (Customer) Login(context.Context, *customer.LoginRequest) (*customer.CustomerMsg, error) {
	return nil, nil
}
