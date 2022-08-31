package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FiberPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MAINPORT")
}

func MONGODB_URI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func Direction() string {
	return MyIp() + FiberPort()
}
