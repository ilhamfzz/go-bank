package service

import (
	"net/smtp"

	"go-bank/domain"
	"go-bank/internal/config"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

func (e emailService) SendEmailVerification(to string, otp string) error {

	auth := smtp.PlainAuth("", e.cnf.Email.Username, e.cnf.Email.Password, e.cnf.Email.Host)
	msg := []byte("" +
		"From: github.com/ilhamfzz <" + e.cnf.Email.Username + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + "Email Verification" + "\r\n" +
		"\r\n" + "Your OTP is " + otp + "\r\n")

	return smtp.SendMail(e.cnf.Email.Host+":"+e.cnf.Email.Port, auth, e.cnf.Email.Username, []string{to}, msg)
}
