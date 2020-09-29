package customer

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Get(ctx context.Context, _ *empty.Empty) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId == 0 {
		return nil, status.Error(codes.Unauthenticated, "session has no customer")
	}

	resp := &customer.CustomerMsg{}
	row := c.gatewayDB.Select(fieldEmail, fieldName).
		From(objectCustomer).
		Where(squirrel.Eq{
			fieldId: session.CustomerId,
		})

	if err := row.Scan(&resp.Email, &resp.Name); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
