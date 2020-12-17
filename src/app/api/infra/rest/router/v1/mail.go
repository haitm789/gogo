package v1

import (
	"api/infra/rest/router/domain"
	"api/interface/repo"
	"api/interface/rest/handler"
	"api/usecase"

	fiber "github.com/gofiber/fiber/v2"
)

func MailRoutes() []domain.Route {
	r := []domain.Route{
		domain.Route{
			Method: "GET",
			URL:    "/test2",
			Handler: func(c *fiber.Ctx) error {
				rp := repo.NewMail()
				uc := usecase.NewMail(rp)
				h := handler.NewMail(uc)
				return h.List(c)
			},
		},
	}

	return r
}
