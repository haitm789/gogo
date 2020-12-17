package repo

import (
	"api/domain"
	"time"
)

type MailRepo struct {
}

func NewMail() MailRepo {
	return MailRepo{}
}

func (repo MailRepo) GetByID(id int) (domain.Mail, error) {
	mailTemplateId := 999
	userId := 998
	params := "params"
	status := 1
	sendAt := time.Now()
	doneAt := time.Now()

	m := domain.NewMail(
		id,
		mailTemplateId,
		userId,
		params,
		status,
		sendAt,
		doneAt,
	)
	return m, nil
}
