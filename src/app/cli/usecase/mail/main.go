package mail

import (
	repository "cli/repo"
)

type Mail struct {
	repo repository.Mail
}

func NewMail(repo repository.Mail) Mail {
	return Mail{
		repo: repo,
	}
}
