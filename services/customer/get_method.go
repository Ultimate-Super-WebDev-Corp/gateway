package customer

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Get(ctx context.Context, _ *empty.Empty) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId == 0 {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	resp := &customer.CustomerMsg{}
	row := c.gatewayDB.Select(fieldEmail, fieldName).
		From(objectCustomer).
		Where(squirrel.Eq{
			fieldId: session.CustomerId,
		})

	if err := row.Scan(&resp.Email, &resp.Name); err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}
	return resp, nil
}
