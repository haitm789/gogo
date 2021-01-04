package database

import (
	"cli/domain"

	"time"
)

type (
	Mail struct {
		ID             int       `gorm:"primary_key;column:id"`
		MailTemplateID int       `gorm:"column:mail_template_id"`
		UserID         int       `gorm:"column:user_id"`
		Params         string    `gorm:"column:params"`
		Status         int       `gorm:"column:status"`
		SendAt         time.Time `gorm:"column:send_at"`
		DoneAt         time.Time `gorm:"column:done_at"`
	}
)

func ToDomainMail(r Mail) domain.Mail {
	return domain.NewMail(
		r.ID,
		r.MailTemplateID,
		r.UserID,
		r.Params,
		r.Status,
		r.SendAt,
		r.DoneAt,
	)
}
