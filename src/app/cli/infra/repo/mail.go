package repo

import (
	"cli/adapter"
	"cli/adapter/database"
	"cli/domain"
)

type (
	Mail struct {
		table string
		db    adapter.Database
	}
)

func NewMail(db adapter.Database) Mail {
	return Mail{
		table: "mails",
		db:    db,
	}
}

func (r Mail) Find() ([]domain.Mail, error) {
	res := []domain.Mail{}

	rs := []database.Mail{}
	r.db.
		Connection().
		Table(r.table).
		Where("status = ?", 1).
		Find(&rs)

	for _, v := range rs {
		res = append(res, database.ToDomainMail(v))
	}

	return res, nil
}
