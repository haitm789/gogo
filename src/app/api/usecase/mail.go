package usecase

import (
	"api/domain"
)

type Mail struct {
	repo MailRepo
}

func NewMail(repo MailRepo) Mail {
	return Mail{
		repo: repo,
	}
}

func (u *Mail) Test() interface{} {
	r := map[string]interface{}{
		"success": 7,
		"message": "test",
	}

	return r
}

func (u *Mail) Get(id int) domain.Mail {
	r2, _ := u.repo.GetByID(id)
	return r2
}
