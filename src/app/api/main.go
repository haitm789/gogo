package main

import (
	"api/infra/rest"
	"api/lib"
	"fmt"
)

func main() {
	port := lib.Config("APP_PORT")
	rest.Start(port)

	fmt.Println("app listen on port", port)
}
