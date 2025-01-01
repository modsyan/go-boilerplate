package email

import (
	"fmt"
	"net/smtp"
)

type IEmailService interface {
	SendVerificationEmail(email, name, verificationLink string) error
	sendEmail(to, subject, body string) error
}

type Service struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewEmailService(host, port, username, password, from string) IEmailService {
	return &Service{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *Service) SendVerificationEmail(email, name, verificationLink string) error {
	subject := "Email Verification"
	body := fmt.Sprintf("Hello %s,\n\nPlease verify your email by clicking the link below:\n%s\n\nThank you.", name, verificationLink)
	return s.sendEmail(email, subject, body)
}

func (s *Service) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}
