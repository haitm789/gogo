package router

import (
	"api/infra/rest/router/domain"
	v1 "api/infra/rest/router/v1"

	fiber "github.com/gofiber/fiber/v2"
)

func Apply(app *fiber.App) {
	gr := []domain.Group{}

	gr = append(gr, v1.Routes())

	for _, g := range gr {
		r := app.Group(g.Prefix)

		for _, rt := range g.Routes {
			switch rt.Method {
			case "GET":
				r.Get(rt.URL, rt.Handler)
			default:
			}
		}
	}
}
