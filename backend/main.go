package main

import (
	"fmt"
	"os"

	"github.com/velicanercan/simple-user-mgmt/infrastructure"
)

func main() {
	infrastructure.NewGinRouter()
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "mysql":
		infrastructure.NewMySQLDatabase()
		fmt.Println("Connected to MySQL database")
	case "mongodb":
	default:
		fmt.Println("No database connection")
	}

}

func init() {
	infrastructure.LoadEnv()
	infrastructure.InitializeLogger()
}
