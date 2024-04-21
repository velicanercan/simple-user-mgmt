package repository

import (
	"context"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/errors"
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (r *UserRepositoryMock) GetUser(ctx context.Context, id int) (*domain.User, error) {
	if id != 1 {
		return nil, errors.ErrUserNotFound
	}
	// return a user with the given id
	return &domain.User{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		BirthDate: "1990-01-01",
	}, nil
}

func (r *UserRepositoryMock) InsertUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *UserRepositoryMock) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return []domain.User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			BirthDate: "1990-01-01",
		},
	}, nil
}

func (r *UserRepositoryMock) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	// validate the user
	if user.Email == "invalidmail" {
		return errors.ErrInvalidEmail
	}
	return nil
}

func (r *UserRepositoryMock) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (r *UserRepositoryMock) Close(ctx context.Context) {}
