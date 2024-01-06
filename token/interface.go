package token

import "time"

type TokenMaker interface {
	CreateToken(string, time.Duration) (string, error)
	VerifeToken(string) (*Payload, error)
}
