package customer

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) Login(ctx context.Context, msg *customer.LoginRequest) (*customer.CustomerMsg, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId != 0 {
		return nil, status.Error(codes.PermissionDenied, "the session has a customer")
	}

	customerId := int64(0)
	password := ""
	resp := &customer.CustomerMsg{}

	row := c.gatewayDB.Select(fieldId, fieldEmail, fieldName, fieldPassword).
		From(objectCustomer).
		Where(squirrel.Eq{
			fieldEmail: msg.Email,
		})

	if err := row.Scan(&customerId, &resp.Email, &resp.Name, &password); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(msg.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, status.Error(codes.NotFound, "wrong password")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	session.CustomerId = customerId
	return resp, nil
}
