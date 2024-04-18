package repository

import (
	"github.com/velicanercan/simple-user-mgmt/infrastructure"
	"github.com/velicanercan/simple-user-mgmt/models"
)

// User Repository is a layer between the database and the service
type UserRepository struct {
	db infrastructure.DBClient
}

// NewUserRepository returns a UserRepository instance
func NewUserRepository(db infrastructure.DBClient) UserRepository {
	return UserRepository{
		db: db,
	}
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user models.User) error {
	resultErr := r.db.InsertUser(&user)
	if resultErr != nil {
		return resultErr
	}
	return nil
}

// GetUserByID returns a user by id
func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	user, err := r.db.GetUser(id)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

// GetAllUsers returns all users
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	users, err := r.db.GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

// UpdateUser updates a user
func (r *UserRepository) UpdateUser(id int, user models.User) error {
	err := r.db.UpdateUser(id, &user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func (r *UserRepository) DeleteUser(id int) error {
	err := r.db.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
