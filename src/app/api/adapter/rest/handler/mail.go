package handler

import (
	"api/adapter/rest"
	"api/domain"
	"api/usecase"

	"time"
)

type (
	Mail struct {
		uc usecase.Mail
	}

	listMailResponse struct {
		ID     int               `json:"id"`
		Name   string            `json:"name"`
		Status domain.MailStatus `json:"status"`
		SendAt time.Time         `json:"send_at"`
		DoneAt time.Time         `json:"done_at"`
	}
)

func NewMail(uc usecase.Mail) Mail {
	return Mail{
		uc: uc,
	}
}

func (m *Mail) List(c rest.Context) error {
	r := m.uc.Get(9)
	res := listMailResponse{
		ID:     r.GetID(),
		Name:   "testing",
		Status: r.GetStatus(),
		SendAt: r.GetSendAt(),
		DoneAt: r.GetDoneAt(),
	}
	return c.Status(200).JSON(&res)
}
