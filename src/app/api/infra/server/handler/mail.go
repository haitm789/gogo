package handler

import (
	"github.com/gofiber/fiber"
)

type Mail struct {

}

func NewMail() Mail {
  return Mail{};
}

func (h Mail) List(c *fiber.Ctx) {
  // c.Send("All Books")
  c.Status(500).JSON(&fiber.Map{
		"success": false,
		"message": "Error creating product",
	})
	// return
}