package util

import (
	"go-wallet.in/dto"
)

func BuildResponse(message string, data any) dto.Response {
	return dto.Response{
		Message: message,
		Errors:  dto.EmptyObj{},
		Data:    data,
	}
}

func BuildErrorResponse(message string, err error) dto.Response {
	return dto.Response{
		Message: message,
		Errors:  err.Error(),
		Data:    dto.EmptyObj{},
	}
}
