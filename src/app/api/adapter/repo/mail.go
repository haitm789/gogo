package repo

import (
	"api/adapter"
	"api/domain"
	"time"
)

type MailRepo struct {
	table string
	db adapter.Database
}

func NewMail(db adapter.Database) MailRepo {
	return MailRepo{
		table: "mails"
		db: db,
	}
}

func (r *MailRepo) GetByID(id int) (domain.Mail, error) {
	res := []domain.Mail{}

	r.db.
		Connection().
		Table(r.table).
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
