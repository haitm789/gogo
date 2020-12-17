package rest

import (
	"api/infra/rest/router"

	fiber "github.com/gofiber/fiber/v2"
)

func Start(port string) {
	app := fiber.New()

	router.Apply(app)

	app.Listen(":" + port)
}
