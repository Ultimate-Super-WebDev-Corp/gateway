package customer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Logout(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	session.CustomerId = 0
	return &empty.Empty{}, nil
}
