package router

import (
	"api/adapter"
	"api/infra/database"
	"api/infra/rest/router/domain"
	v1 "api/infra/rest/router/v1"

	fiber "github.com/gofiber/fiber/v2"
)

func Apply(app *fiber.App) {
	gr := []domain.Group{}

	db := db()
	gr = append(gr, v1.Routes(db))

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

func db() adapter.Database {
	return database.NewMySql()
}
