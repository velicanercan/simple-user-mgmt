package infrastructure

import (
	"os"

	"github.com/velicanercan/simple-user-mgmt/models"
)

type DBClient interface {
	// InsertUser inserts a user into the database
	InsertUser(user *models.User) error
	// GetAllUsers retrieves all users from the database
	GetAllUsers() ([]models.User, error)
	// GetUser retrieves a user from the database by ID
	GetUser(id int) (*models.User, error)
	// UpdateUser updates a user in the database
	UpdateUser(id int, user *models.User) error
	// DeleteUser deletes a user from the database by ID
	DeleteUser(id int) error
	// Close closes the database connection
	Close()
}

type Database struct {
	DB DBClient
}

// InitializeDB initializes the database based on the DB_TYPE environment variable
func InitializeDB() DBClient {
	// get the database type from the environment variable
	DB_TYPE := os.Getenv("DB_TYPE")

	// initialize the database based on the type
	switch DB_TYPE {
	case "mysql":
		return NewMySQLDatabase()
	case "mongodb":
		return NewMongoDBDatabase()
	default:
		panic("Invalid database type!")
	}
}
