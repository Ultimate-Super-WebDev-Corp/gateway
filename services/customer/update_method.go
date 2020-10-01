package customer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Update(ctx context.Context, msg *customer.UpdateRequest) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId == 0 {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	updBuilder := c.gatewayDB.
		Update(objectCustomer).
		Where(squirrel.Eq{
			fieldId:         session.CustomerId,
			fieldPasswordId: session.PasswordId,
		}).
		Suffix(fmt.Sprintf("returning %s, %s", fieldName, fieldEmail))

	wasSet := false
	if len(msg.Name) > 0 {
		wasSet = true
		updBuilder = updBuilder.Set(fieldName, msg.Name)
	}

	if !wasSet {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("no updates"))
	}

	query := updBuilder.QueryRow()

	resp := customer.CustomerMsg{}
	if err := query.Scan(&resp.Name, &resp.Email); err != nil {
		if err == sql.ErrNoRows {
			server.SessionLogout(session)
			return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &resp, nil
}
