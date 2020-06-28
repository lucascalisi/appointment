package utils

import (
	"fmt"
	"net/smtp"
)

type EmailSender struct {
	User             string
	MySecretPassword string
	SMTPServer       string
	Port             string
}

func NewEmailSender(user string, password string, server string, port string) (EmailSender, error) {
	return EmailSender{
		User:             user,
		MySecretPassword: password,
		SMTPServer:       server,
		Port:             port,
	}, nil
}

func (e EmailSender) SendCancelAppointmentNotificacion(addressee string, subject string, body string) error {
	fmt.Println(e)
	from := e.User
	password := e.MySecretPassword
	// Receiver email address.
	to := []string{
		addressee,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: e.SMTPServer, port: e.Port}
	// Message.
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
