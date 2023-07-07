package domain

type EmailService interface {
	SendEmailVerification(to string, otp string) error
}
