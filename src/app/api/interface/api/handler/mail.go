package handler

import (
	"api/usecase"

	fiber "github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	uc := usecase.NewMail()
	r := uc.List()
	return c.Status(200).JSON(&r)
}
