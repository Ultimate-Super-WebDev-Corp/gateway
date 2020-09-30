package customer

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Login(ctx context.Context, msg *customer.LoginRequest) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId != 0 {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("the session has a customer"))
	}

	customerId := int64(0)
	passwordId := int64(0)
	password := ""
	resp := &customer.CustomerMsg{}

	row := c.gatewayDB.Select(fieldId, fieldEmail, fieldName, fieldPassword, fieldPasswordId).
		From(objectCustomer).
		Where(squirrel.Eq{
			fieldEmail: msg.Email,
		})

	if err := row.Scan(&customerId, &resp.Email, &resp.Name, &password, &passwordId); err != nil {
		if err == sql.ErrNoRows {
			return nil, server.NewErrServer(codes.NotFound, errors.New("customer not found"))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	if err := comparePassword(password, msg.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, server.NewErrServer(codes.NotFound, errors.New("wrong password"))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	server.SessionLogin(session, customerId, passwordId)
	return resp, nil
}
