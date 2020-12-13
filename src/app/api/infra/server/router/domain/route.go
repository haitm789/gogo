package domain

import (
  "github.com/gofiber/fiber/v2"
)

type (
  Route struct {
    Method     string
    URL        string
    Handler   fiber.Handler
    Middlewares []fiber.Handler
  }
)
