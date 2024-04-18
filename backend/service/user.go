package service

import (
	"time"

	"github.com/velicanercan/simple-user-mgmt/errors"
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
	// check if the user already exists
	_, err := s.repository.GetUserByID(user.ID)
	if err == nil {
		return errors.ErrUserAlreadyExists
	}
	// convert the birthdate string to a time.Time object
	birthdate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return errors.ErrInvalidBirthDate
	}
	// check user age is greater than 18
	age := time.Since(birthdate).Hours() / 24 / 365
	if age < 18 {
		return errors.ErrUserAge
	}
	return s.repository.CreateUser(user)
}

// GetUserByID returns a user by id
func (s *UserService) GetUserByID(id int) (models.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return models.User{}, errors.ErrUserNotFound
	}
	return user, nil
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id int, user models.User) error {
	// check if the user already exists
	_, err := s.repository.GetUserByID(id)
	if err != nil {
		return errors.ErrUserAlreadyExists
	}
	// convert the birthdate string to a time.Time object
	birthdate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return errors.ErrInvalidBirthDate
	}
	// check user age is greater than 18
	age := time.Since(birthdate).Hours() / 24 / 365
	if age < 18 {
		return errors.ErrUserAge
	}
	return s.repository.UpdateUser(id, user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}
