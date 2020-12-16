package server

import (
	"api/infra/server/router"

	fiber "github.com/gofiber/fiber/v2"
)

func Start(port string) {
	app := fiber.New()

	router.Apply(app)

	app.Listen(":" + port)
}
