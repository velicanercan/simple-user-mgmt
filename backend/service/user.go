package service

import (
	"github.com/velicanercan/simple-user-mgmt/models"
	"github.com/velicanercan/simple-user-mgmt/repository"
)

// UserService is a layer between the repository and the controller
type UserService struct {
	repository repository.UserRepository
}

// NewUserService returns a UserService instance
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user models.User) error {
	return s.repository.CreateUser(user)
}

// GetUserByID returns a user by id
func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repository.GetUserByID(id)
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(user models.User) error {
	return s.repository.UpdateUser(user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}
