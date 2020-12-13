package server

import (
   "api/infra/server/router"
  "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/requestid"
)

func routes (app *fiber.App) { 
  // Middleware
  // api := app.Group("/api", logger.New(), middleware.AuthReq())  
  app.Use(requestid.New())
  r := router.NewRouter()
  r.Apply(app)
}

func Start(port string) {
  app := fiber.New()

  routes(app)

  app.Listen(":" + port)
}
