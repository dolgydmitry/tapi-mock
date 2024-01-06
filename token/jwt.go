package token

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenType     = "RS256"
	ExpireTimeMin = 1
)

type JwtMaker struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (m *JwtMaker) CreateToken(user string, duration time.Duration) (string, error) {
	payload, err := NewPayload(user, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.New(jwt.GetSigningMethod(TokenType))
	jwtToken.Claims = payload
	return jwtToken.SignedString(m.PrivateKey)

}

func (m *JwtMaker) VerifeToken(accessToken string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return m.PublicKey, nil
	}
	jwtToken, err := jwt.ParseWithClaims(accessToken, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil

}
