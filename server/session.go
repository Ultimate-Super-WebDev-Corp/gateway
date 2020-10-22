package server

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"github.com/ulule/deepcopier"
	"go.uber.org/zap"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/general"
)

const mdToken = "token"

type sessionClaims struct {
	jwt.StandardClaims
	general.Session
}

const sessionTTL = 15 * time.Minute

func (s Server) getSession(md map[string][]string) (*general.Session, error) {
	strToken := md[mdToken]
	if len(strToken) == 0 {
		return nil, errors.New("authorization token is not provided")
	}

	accessToken := strToken[0]
	token, err := jwt.ParseWithClaims(accessToken, &sessionClaims{}, s.keyFunc)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*sessionClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	updatedAt := time.Unix(0, claims.UpdatedAt)
	if updatedAt.Add(sessionTTL).Before(time.Now().UTC()) {
		return nil, errors.New("session is too old")
	}
	if claims.Session.Id == "" {
		UUID, err := uuid.NewUUID()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		claims.Session.Id = UUID.String()
	}

	return &claims.Session, nil
}

func (s Server) makeSessionToken(ctx context.Context) (string, error) {
	session := SessionFromCtx(ctx)
	session.UpdatedAt = time.Now().UTC().UnixNano()
	respToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims{Session: *session})

	strRespToken, err := respToken.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return strRespToken, nil
}

var ctxSessionMarkerKey = &ctxSessionMarker{}

type ctxSessionMarker struct{}

func sessionToCtx(ctx context.Context, session *general.Session) context.Context {
	ctxzap.AddFields(ctx,
		zap.String("session_id", session.Id),
		zap.Int64("customer_id", session.CustomerId))

	return context.WithValue(ctx, ctxSessionMarkerKey, session)
}

func SessionInCtxUpdate(ctx context.Context, newSession *general.Session) {
	session := SessionFromCtx(ctx)
	_ = deepcopier.Copy(newSession).To(session)
	session.UpdatedAt = time.Now().UTC().UnixNano()
}

func SessionFromCtx(ctx context.Context) *general.Session {
	session, ok := ctx.Value(ctxSessionMarkerKey).(*general.Session)
	if !ok || session == nil {
		return &general.Session{} //todo return error?
	}
	return session
}

func IsSessionLoggedIn(s *general.Session) bool {
	return s.CustomerId > 0
}
func IsSessionRoot(s *general.Session) bool {
	return s.CustomerId > 0 && s.CustomerId < 1000 && s.PasswordId == -1

}
func SessionLogout(s *general.Session) {
	s.CustomerId = 0
	s.PasswordId = 0
}

func SessionLogin(s *general.Session, cusId int64, passId int64) {
	s.CustomerId = cusId
	s.PasswordId = passId
}
