package dto

type Response struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}
