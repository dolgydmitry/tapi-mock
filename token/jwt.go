package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenType     = "RS256"
	ExpireTimeMin = 1
)

type TokenMaker interface {
	CreateToken(user string) (string, error)
	VerifeToken()
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	CustomerInfo
}

type CustomerInfo struct {
	Name string
	Kind string
}

type JwtMaker struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (m *JwtMaker) CreateToken(user string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(TokenType))
	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * ExpireTimeMin).Unix(),
		},
		CustomerInfo{user, "human"},
	}
	return t.SignedString(m.PrivateKey)
}

func (m *JwtMaker) VerifeToken() {}
