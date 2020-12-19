package rest

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Context interface {
	Params(string, ...string) string
	Status(int) *fiber.Ctx
	JSON(interface{}) error
}
