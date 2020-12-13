package router

import (
  "api/infra/server/router/domain"
  "api/infra/server/router/v1"

  "github.com/gofiber/fiber/v2"
)

type (
  router struct {}
)

func NewRouter() router {
  return router{}
}

func (s *router) Apply(app *fiber.App) {
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
