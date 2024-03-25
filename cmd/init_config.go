package cmd

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Could not load .env file.")
	}
}
