package v1

import (
	"api/infra/rest/router/domain"
)

func user() []domain.Route {
	r := []domain.Route{
		// domain.Route{
		// 	Method: "GET",
		// 	URL:    "/user/:name",
		// 	Handler: func(c *fiber.Ctx) error {
		// 		rp := repo.NewMail()
		// 		uc := usecase.NewMail(rp)
		// 		h := handler.NewUser(uc)
		// 		return h.User(c)
		// 	},
		// },
	}

	return r
}
