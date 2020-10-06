package customer

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Get(ctx context.Context, _ *empty.Empty) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if !server.IsSessionLoggedIn(session) {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	resp := &customer.CustomerMsg{}
	row := c.statementBuilder.Select(fieldEmail, fieldName).
		From(objectCustomer).
		Where(squirrel.Eq{
			fieldId:         session.CustomerId,
			fieldPasswordId: session.PasswordId,
		}).RunWith(c.gatewayDB).QueryRow()

	if err := row.Scan(&resp.Email, &resp.Name); err != nil {
		if err == sql.ErrNoRows {
			server.SessionLogout(session)
			return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}
	return resp, nil
}
