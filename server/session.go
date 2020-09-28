package server

import (
	"context"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type sessionClaims struct {
	jwt.StandardClaims
	model.Session
}

func (s Server) getSession(md map[string][]string) (*model.Session, error) {
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
	respToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims{Session: *session})

	strRespToken, err := respToken.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return strRespToken, nil
}

var ctxSessionMarkerKey = &ctxSessionMarker{}

type ctxSessionMarker struct{}

func sessionToCtx(ctx context.Context, session *model.Session) context.Context {
	return context.WithValue(ctx, ctxSessionMarkerKey, session)
}

func SessionInCtxUpdate(ctx context.Context, newSession *model.Session) {
	session := SessionFromCtx(ctx)
	session.Id = newSession.Id
}
func SessionFromCtx(ctx context.Context) *model.Session {
	session, ok := ctx.Value(ctxSessionMarkerKey).(*model.Session)
	if !ok || session == nil {
		return &model.Session{} //todo return error?
	}
	return session
}
