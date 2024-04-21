package repository

import (
	"context"
	"os"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository/mongo"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository/mysql"
)

type UserRepository interface {
	// InsertUser inserts a user into the database
	InsertUser(ctx context.Context, user *domain.User) error
	// GetAllUsers retrieves all users from the database
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	// GetUser retrieves a user from the database by ID
	GetUser(ctx context.Context, id int) (*domain.User, error)
	// UpdateUser updates a user in the database
	UpdateUser(ctx context.Context, id int, user *domain.User) error
	// DeleteUser deletes a user from the database by ID
	DeleteUser(ctx context.Context, id int) error
	// Close closes the database connection
	Close(ctx context.Context)
}

type Database struct {
	DB UserRepository
}

// InitializeDB initializes the database based on the DB_TYPE environment variable
func InitializeDB(ctx context.Context) UserRepository {
	// get the database type from the environment variable
	DB_TYPE := os.Getenv("DB_TYPE")

	// initialize the database based on the type
	switch DB_TYPE {
	case "mysql":
		return mysql.NewMySQLDatabase()
	case "mongodb":
		return mongo.NewMongoDBDatabase(ctx)
	default:
		panic("Invalid database type!")
	}
}
