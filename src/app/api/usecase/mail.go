package usecase

import (
	"api/domain"
	repository "api/repo"
)

type Mail struct {
	repo repository.Mail
}

func NewMail(repo repository.Mail) Mail {
	return Mail{
		repo: repo,
	}
}

func (u *Mail) Get(id int) domain.Mail {
	r, _ := u.repo.GetByID(id)
	return r
}
