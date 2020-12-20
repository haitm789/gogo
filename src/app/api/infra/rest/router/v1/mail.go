package v1

import (
	"api/adapter"
	"api/adapter/repo"
	"api/adapter/rest/handler"
	"api/infra/rest/router/domain"
	"api/usecase"

	fiber "github.com/gofiber/fiber/v2"
)

func mail(db adapter.Database) []domain.Route {
	r := []domain.Route{
		domain.Route{
			Method: "GET",
			URL:    "/mail",
			Handler: func(c *fiber.Ctx) error {
				rp := repo.NewMail(db)
				uc := usecase.NewMail(rp)
				h := handler.NewMail(uc)
				return h.List(c)
			},
		},
	}

	return r
}
