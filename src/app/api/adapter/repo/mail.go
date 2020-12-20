package repo

import (
	"api/adapter"
	"api/domain"
	"time"
)

type MailRepo struct {
	db adapter.Database
}

func NewMail(db adapter.Database) MailRepo {
	return MailRepo{
		db: db,
	}
}

func (repo MailRepo) GetByID(id int) (domain.Mail, error) {
	res := []domain.Mail{}

	repo.db.
		Connection().
		Table("mails").
		Where("send_at <= NOW() AND status = ?", 1).
		Find(&res)

	mailTemplateId := 997
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
