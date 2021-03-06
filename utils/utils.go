package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Init to load .env file
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
