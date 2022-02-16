package controllers

import (
	"fmt"
	"net/smtp"
	"os"
)

type Mail struct {
	From     string
	Auth     smtp.Auth
	Username string
	To       []string
	Message  []byte
}

func SendMail(body string) error {
	var mail Mail
	mail.Auth = smtp.PlainAuth(
		"",
		os.Getenv("SMTP_EMAIL_USERNAME"),
		os.Getenv("SMTP_EMAIL_PASSWORD"),
		"smtp.gmail.com",
	)
	mail.To = []string{
		"ivankonewv@gmail.com",
	}
	mail.Username = os.Getenv("SMTP_EMAIL_USERNAME")
	mail.From = os.Getenv("SMTP_PROVIDER_ADDRESS")

	mail.Message = []byte(fmt.Sprintf("To: %s\nSubject: Новая заявка на звонок на сайте.\n%s", mail.To, body))

	if err := smtp.SendMail(mail.From, mail.Auth, mail.Username, mail.To, mail.Message); err != nil {
		return fmt.Errorf("cannot send message: %v", err)
	}

	return nil
}
