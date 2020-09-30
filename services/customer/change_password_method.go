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

func (c Customer) ChangePassword(ctx context.Context, msg *customer.ChangePasswordRequest) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	if session.CustomerId == 0 {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	password, err := generatePassword(msg.NewPassword)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	res, err := c.gatewayDB.
		Update(objectCustomer).
		Set(fieldPassword, password).
		Where(squirrel.Eq{
			fieldId:         session.CustomerId,
			fieldPasswordId: session.PasswordId,
		}).Exec()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}
	if affected == 0 {
		server.SessionLogout(session)
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	return &empty.Empty{}, nil
}
