package dto

type ValidateOTPReq struct {
	RefrenceID string `json:"refrence_id"`
	OTP        string `json:"otp"`
}
