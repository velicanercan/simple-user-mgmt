package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/velicanercan/simple-user-mgmt/logger"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Log("panic", "Failed to load environment variables!", err.Error())
	}
}
