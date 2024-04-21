package service

import (
	"context"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository"
)

type UserServiceMock struct {
	repository repository.UserRepository
}

func NewUserServiceMock(repository repository.UserRepository) UserServiceMock {
	return UserServiceMock{
		repository: repository,
	}
}

func (uc *UserServiceMock) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}

func (uc *UserServiceMock) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return domain.User{}, nil
}

func (uc *UserServiceMock) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (uc *UserServiceMock) UpdateUser(ctx context.Context, id int, user domain.User) error {
	return nil
}

func (uc *UserServiceMock) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (uc *UserServiceMock) CheckUserEmail(ctx context.Context, user domain.User, id int) error {
	return nil
}

func (uc *UserServiceMock) On(method string, args ...interface{}) {
	switch method {
	case "CreateUser":
		uc.CreateUser(args[0].(context.Context), args[1].(domain.User))
	case "GetUserByID":
		uc.GetUserByID(args[0].(context.Context), args[1].(int))
	case "GetAllUsers":
		uc.GetAllUsers(args[0].(context.Context))
	case "UpdateUser":
		uc.UpdateUser(args[0].(context.Context), args[1].(int), args[2].(domain.User))
	case "DeleteUser":
		uc.DeleteUser(args[0].(context.Context), args[1].(int))
	case "CheckUserEmail":
		uc.CheckUserEmail(args[0].(context.Context), args[1].(domain.User), args[2].(int))
	}
}