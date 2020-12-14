package handler

import (
	"api/usecase"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	uc := usecase.NewMail()
	r:= uc.Get()
	return c.Status(200).JSON(&r)
}