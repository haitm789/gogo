package usecase

import (
	repository "cli/adapter/repo"
	"fmt"
)

type Mail struct {
	repo repository.Mail
}

func NewMail(repo repository.Mail) Mail {
	return Mail{
		repo: repo,
	}
}
func (u *Mail) Send(str string) {
	fmt.Println("slack gateway: " + str)
}
