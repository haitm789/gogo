package v1

import (
	"api/infra/server/router/domain"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes() ([]domain.Route) {
  r := []domain.Route{
		domain.Route{
			Method:  "GET",
			URL: "/posts/test1",
			Handler:    func(c *fiber.Ctx) error {
        return c.SendString("posts/test1")
      },
			// Middlewares: []gin.HandlerFunc{},
		},
		domain.Route{
			Method:  "GET",
			URL: "/posts/test2",
			Handler:    func(c *fiber.Ctx) error {
        return c.SendString("posts/test1")
      },
			// Middlewares: []gin.HandlerFunc{},
		},
	}
  
  return r;
} 
