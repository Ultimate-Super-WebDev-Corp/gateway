package customer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (c Customer) ChangePassword(ctx context.Context, msg *customer.ChangePasswordRequest) (*empty.Empty, error) {
	session := server.SessionFromCtx(ctx)
	if !server.IsSessionLoggedIn(session) {
		return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
	}

	password, err := generatePassword(msg.NewPassword)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	res := c.statementBuilder.
		Update(objectCustomer).
		Set(fieldPassword, password).
		Set(fieldPasswordId, squirrel.ConcatExpr(fieldPasswordId, " + 1")).
		Where(squirrel.Eq{
			fieldId:         session.CustomerId,
			fieldPasswordId: session.PasswordId,
		}).
		Suffix(fmt.Sprintf("returning %s", fieldPasswordId)).
		RunWith(c.gatewayDB).QueryRow()

	passwordId := int64(0)
	if err := res.Scan(&passwordId); err != nil {
		if err == sql.ErrNoRows {
			server.SessionLogout(session)
			return nil, server.NewErrServer(codes.Unauthenticated, errors.New("session has no customer"))
		}
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	server.SessionLogin(session, session.CustomerId, passwordId)
	return &empty.Empty{}, nil
}
