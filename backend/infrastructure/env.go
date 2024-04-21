package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/velicanercan/simple-user-mgmt/logger"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Log("error", "Error loading .env file, using default values...")
		loadDefaultEnv()
	}
}

func loadDefaultEnv() {
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_TYPE", "mysql")
	os.Setenv("MYSQL_ROOT_PASSWORD", "root")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "verysecret")
	os.Setenv("DB_HOST", "mysql")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "user")
	os.Setenv("ENVIRONMENT", "local")
	os.Setenv("MONGO_URI", "mongodb://mongo:27017")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_DB_NAME", "user")
	os.Setenv("CORS_POLICY", "*")
}