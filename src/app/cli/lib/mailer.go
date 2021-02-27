package lib

import (
	"cli/adapter/mail"
)

type (
	Mailer struct {
		mailer mail.Mailer
	}
)

func NewMailer(
	mailer mail.Mailer,
) Mailer {
	return Mailer{
		mailer: mailer,
	}
}

func (m Mailer) Flush(mailw mail.Mail) {
	m.mailer.Flush(mailw)
}
