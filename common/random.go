package common

import (
	"errors"
	"github.com/google/uuid"
)

func GenUUID() (string, error) {
	uuidRand, err := uuid.NewRandom()
	if err != nil {
		return "", ErrGenerateUUID
	}
	return uuidRand.String(), nil
}

var (
	ErrGenerateUUID = NewCustomError(errors.New("cannot generate uuid"),
		"cannot generate uuid",
		"ErrCannotGenerateUUID")
)
