package usecase

import (
	"fmt"
)

type (
	mail struct{}
)

func NewMail() mail {
	return mail{}
}

func (m *mail) List() map[string]string {
	fmt.Println("uc")
	r := map[string]string{
		"message": "ok",
	}
	return r
}
