package handler

import (
	"api/domain"
	"api/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	Mail struct {
		uc usecase.Mail
	}

	Response struct {
		ID     int               `json:"id"`
		Status domain.MailStatus `json:"status"`
		SendAt time.Time         `json:"send_at"`
		DoneAt time.Time         `json:"done_at"`
	}
)

func NewMail(uc usecase.Mail) *Mail {
	return &Mail{
		uc: uc,
	}
}

func (m *Mail) List(c *fiber.Ctx) error {
	r := m.uc.Get(9)

	res := Response{
		ID:     r.GetID(),
		Status: r.GetStatus(),
		SendAt: r.GetSendAt(),
		DoneAt: r.GetDoneAt(),
	}
	return c.Status(200).JSON(&res)
}
