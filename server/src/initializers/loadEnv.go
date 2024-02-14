package initializers

import (
	"github.com/joho/godotenv"
	"log"

	_ "github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
