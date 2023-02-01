package tokenprovider

import (
	"errors"
	"food_delivery/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(myToken string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound")

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrEncodingToken")

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken")
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

func (t *TokenPayload) GetUserId() int {
	return t.UserId
}

func (t *TokenPayload) GetRole() string {
	return t.Role
}
