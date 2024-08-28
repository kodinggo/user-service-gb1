package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadWithGodotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
