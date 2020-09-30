package customer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Logout(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	server.SessionLogout(server.SessionFromCtx(ctx))
	return &empty.Empty{}, nil
}
