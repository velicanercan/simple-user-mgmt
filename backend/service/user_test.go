package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/errors"
	"github.com/velicanercan/simple-user-mgmt/mocks/repository"
)

func TestUserService_CreateUser(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	userService := NewUserService(repo)

	user := domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		BirthDate: "1990-01-01",
	}

	err := userService.CreateUser(context.Background(), user)
	assert.Equal(t, errors.ErrUserAlreadyExists, err)

	user, err = userService.GetUserByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "John", user.FirstName)

	err = userService.CreateUser(context.Background(), user)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrUserAlreadyExists, err)

	user.Email = "invalidmail"
	err = userService.UpdateUser(context.Background(), user.ID, user)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrInvalidEmail, err)

}

func TestUserService_GetUserByID(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	userService := NewUserService(repo)

	user, err := userService.GetUserByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "John", user.FirstName)

	_, err = userService.GetUserByID(context.Background(), 2)
	assert.Equal(t, errors.ErrUserNotFound, err)
}

func TestUserService_GetAllUsers(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	userService := NewUserService(repo)

	users, err := userService.GetAllUsers(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	userService := NewUserService(repo)

	user := domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		BirthDate: "1990-01-01",
	}

	err := userService.UpdateUser(context.Background(), user.ID, user)
	assert.NoError(t, err)

	user.Email = "invalidmail"
	err = userService.UpdateUser(context.Background(), user.ID, user)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrInvalidEmail, err)
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	userService := NewUserService(repo)

	err := userService.DeleteUser(context.Background(), 1)
	assert.NoError(t, err)
}

func TestCheckUserAge(t *testing.T) {
	err := CheckUserAge("2005-01-01")
	assert.NoError(t, err)

	err = CheckUserAge("2010-01-01")
	assert.Error(t, err)

	err = CheckUserAge("2015-01-01")
	assert.Error(t, err)
}

func TestCheckUserEmail(t *testing.T) {
	user := domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "invalidmail",
		BirthDate: "1990-01-01",
	}
	service := NewUserService(&repository.UserRepositoryMock{})
	err := service.CheckUserEmail(context.Background(), user, user.ID)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrInvalidEmail, err)

	user.Email = "valid@mail.com"
	err = service.CheckUserEmail(context.Background(), user, user.ID)
	assert.NoError(t, err)
}
