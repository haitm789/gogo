package v1

import (
	"api/adapter"
	"api/infra/rest/router/domain"
	"api/wire"

	fiber "github.com/gofiber/fiber/v2"
)

func mail(db adapter.Database) []domain.Route {
	r := []domain.Route{
		domain.Route{
			Method: "GET",
			URL:    "/mail",
			Handler: func(c *fiber.Ctx) error {
				h := wire.InitializeMailHandler(db)
				return h.List(c)
			},
		},
	}

	return r
}
