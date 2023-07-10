package util

import (
	"errors"

	"go-wallet.in/domain"
)

const (
	HttpSatusCreated = 201
	HttpSatusOk      = 200
)

func GetErrHttpStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.ErrOTPInvalid):
		return 400
	case errors.Is(err, domain.ErrOTPExpired):
		return 400
	case errors.Is(err, domain.ErrBadRequest):
		return 400
	case errors.Is(err, domain.ErrUsernameAlreadyExist):
		return 400
	case errors.Is(err, domain.ErrEmailAlreadyExist):
		return 400
	default:
		return 500
	}
}
