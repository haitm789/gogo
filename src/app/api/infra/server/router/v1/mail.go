package v1

import (
	"api/infra/server/router/domain"
	"api/interface/api/handler"

	fiber "github.com/gofiber/fiber/v2"
)

func MailRoutes() []domain.Route {
	r := []domain.Route{
		domain.Route{
			Method: "GET",
			URL:    "/test1",
			Handler: func(c *fiber.Ctx) error {
				return c.SendString("mail/test1")
			},
		},
		domain.Route{
			Method:  "GET",
			URL:     "/test2",
			Handler: handler.List,
		},
	}

	return r
}
