package mail

import (
	a "cli/adapter/mail"
	r "cli/adapter/repo"
)

type Mail struct {
	repo r.Mail
}

func NewMail(repo r.Mail) Mail {
	return Mail{
		repo: repo,
	}
}

func NewMailMessage(
	sender string,
	recipient string,
	subject string,
	body string,
	attachments []string,
	cc []string,
	bcc []string,
) a.Mail {
	return a.NewMail(
		sender,
		recipient,
		subject,
		body,
		attachments,
		cc,
		bcc,
	)
}
