package config

import (
	"log"

	env "github.com/joho/godotenv"
)

func Load() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
