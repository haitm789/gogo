package handler

import (
	"api/domain"
	"api/interface/rest"
	"api/usecase"

	"time"
)

type (
	user struct {
		uc usecase.Mail
	}

	getUserResponse struct {
		ID     int               `json:"id"`
		Name   string            `json:"name"`
		Status domain.MailStatus `json:"status"`
		SendAt time.Time         `json:"send_at"`
		DoneAt time.Time         `json:"done_at"`
	}
)

func NewUser(uc usecase.Mail) *user {
	return &user{
		uc: uc,
	}
}

func (m *user) User(c rest.Context) error {
	name := c.Params("name")

	r := m.uc.Get(10)

	res := getUserResponse{
		ID:     r.GetID(),
		Name:   name,
		Status: r.GetStatus(),
		SendAt: r.GetSendAt(),
		DoneAt: r.GetDoneAt(),
	}
	return c.Status(202).JSON(&res)
}
