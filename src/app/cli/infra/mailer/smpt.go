package mailer

import (
	"bytes"
	"fmt"
	agent "net/smtp"
	"os"

	"cli/adapter/mail"
)

type (
	smtp struct {
	}
)

func NewSmtpMailer() mail.Mailer {
	return &smtp{}
}

func (*smtp) Flush(msg mail.Mail) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	addr := host + ":" + port

	from := msg.Sender
	to := []string{msg.Recipient}

	auth := auth()
	err := agent.SendMail(addr, auth, from, to, makeMsg(msg))
	fmt.Println(err)
	return err
}

func auth() agent.Auth {
	host := os.Getenv("SMTP_HOST")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	auth := agent.PlainAuth("", username, password, host)
	return auth
}

// https://gist.github.com/douglasmakey/90753ecf37ac10c25873825097f46300
func makeMsg(m mail.Mail) []byte {
	buf := bytes.NewBuffer(nil)
	// withAttachments := len(m.Attachments) > 0

	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString("MIME-Version: 1.0\n")
	// writer := multipart.NewWriter(buf)
	// boundary := writer.Boundary()

	// if withAttachments {
	// 	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
	// 	buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	// }

	buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	buf.WriteString(m.Body)

	// if withAttachments {
	// 	for k, v := range m.Attachments {
	// 		buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
	// 		buf.WriteString("Content-Type: application/octet-stream\n")
	// 		buf.WriteString("Content-Transfer-Encoding: base64\n")
	// 		buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

	// 		b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
	// 		base64.StdEncoding.Encode(b, v)
	// 		buf.Write(b)
	// 		buf.WriteString(fmt.Sprintf("\n--%s", boundary))
	// 	}

	// 	buf.WriteString("--")
	// }

	return buf.Bytes()
}
