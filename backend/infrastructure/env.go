package infrastructure

import (
	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		Log("panic", "Failed to load environment variables!", err.Error())
	}
}
