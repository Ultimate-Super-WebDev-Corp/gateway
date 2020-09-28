package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func (s Server) keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("unexpected token signing method")
	}
	return []byte(s.cfg.SecretKey), nil
}
