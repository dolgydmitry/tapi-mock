package token

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("invalid token")
var ErrExpiredToken = errors.New("token has expired")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, expireTime time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Println("error to create new uuid bly: ", err)
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(expireTime),
	}
	return payload, nil
}
