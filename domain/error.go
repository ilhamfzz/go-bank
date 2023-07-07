package domain

import "errors"

var (
	ErrAuthFailed           = errors.New("authentication failed")
	ErrInternalServerError  = errors.New("internal server error")
	ErrOTPInvalid           = errors.New("otp invalid")
	ErrOTPExpired           = errors.New("otp expired")
	ErrBadRequest           = errors.New("bad request")
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrEmailAlreadyExist    = errors.New("email already exist")
	ErrAccountNotFound      = errors.New("account not found")
	ErrInquiryNotFound      = errors.New("inquiry not found")
	ErrInsufficientBalance  = errors.New("insufficient balance")
)
