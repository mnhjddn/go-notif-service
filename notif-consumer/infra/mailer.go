package infra

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type IMailer interface {
	SendEmail(to []string, subject, body string)
}

type Mailer struct {
	config *AppConfig
}

func NewMailer(config *AppConfig) *Mailer {
	return &Mailer{config}
}

func (m *Mailer) SendEmail(to []string, subject, body string) {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", m.config.Email.Sender)

	mailer.SetHeader("To", to...)

	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		m.config.Email.Host,
		m.config.Email.Port,
		m.config.Email.Usr,
		m.config.Email.Pw,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println(err)
		return
	}
}
