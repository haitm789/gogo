package main

import (
	"api/lib"
  "fmt"
  "api/infra/server"
)

func main() {
  port := lib.Config("APP_PORT");
  server.Start(port);

  fmt.Println("app listen on port", port)
}