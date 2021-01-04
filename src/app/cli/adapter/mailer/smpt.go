package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	agent "net/smtp"
)

type (
	smtp struct{}
)

func NewSmtpMailer() smtp {
	return smtp{}
}

func (m smtp) Send() {
	// Sender data.
	from := "vietsolsmtp@gmail.com"
	password := "ngohongd@0"

	// Receiver email address.
	to := []string{
		"haitm789@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := agent.PlainAuth("", from, password, smtpHost)

	// t, _ := template.ParseFiles("template.html")
	t, _ := template.New("msg").Parse("hello world")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	fmt.Println("auth:", auth)
	fmt.Println("to:", to)
	fmt.Println("from:", from)

	// Sending email.
	err := agent.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
