package service

import (
	"context"
	"time"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/errors"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository"
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
func (s *UserService) CreateUser(ctx context.Context, user domain.User) error {
	// check if the user already exists
	_, err := s.repository.GetUser(ctx, user.ID)
	if err == nil {
		return errors.ErrUserAlreadyExists
	}
	// check if the mail already exists
	err = s.CheckUserEmailUnique(ctx, user, user.ID)
	if err != nil {
		return err
	}
	err = CheckUserAge(user.BirthDate)
	if err != nil {
		return err
	}
	return s.repository.InsertUser(ctx, &user)
}

// GetUserByID returns a user by id
func (s *UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	user, err := s.repository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, errors.ErrUserNotFound
	}
	return *user, nil
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repository.GetAllUsers(ctx)
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(ctx context.Context, id int, user domain.User) error {
	// check if the user already exists
	_, err := s.repository.GetUser(ctx, id)
	if err != nil {
		return errors.ErrUserAlreadyExists
	}
	// check if the mail already exists
	err = s.CheckUserEmailUnique(ctx, user, id)
	if err != nil {
		return err
	}
	// check user age is greater than 18
	err = CheckUserAge(user.BirthDate)
	if err != nil {
		return err
	}
	return s.repository.UpdateUser(ctx, id, &user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repository.DeleteUser(ctx, id)
}

// CheckUserAge checks if the user is older than 18
func CheckUserAge(birthdate string) error {
	// convert the birthdate string to a time.Time object
	birthdateObj, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return errors.ErrInvalidBirthDate
	}
	// check user age is greater than 18
	age := time.Since(birthdateObj).Hours() / 24 / 365
	if age < 18 {
		return errors.ErrUserAge
	}
	return nil
}

// CheckUserExists checks if the user already exists
func (s *UserService) CheckUserEmailUnique(ctx context.Context, user domain.User, id int) error {
	// check if the mail already exists
	users, err := s.repository.GetAllUsers(ctx)
	if err == nil {
		for _, u := range users {
			if u.Email == user.Email && u.ID != id {
				return errors.ErrUserMailExists
			}
		}
	}
	return nil
}
