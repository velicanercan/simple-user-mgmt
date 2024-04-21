package controller

import (
	"context"

	"github.com/velicanercan/simple-user-mgmt/domain"
)

type UserControllerMock struct {
}

func NewUserControllerMock() *UserControllerMock {
	return &UserControllerMock{}
}

func (r *UserControllerMock) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}

func (r *UserControllerMock) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserControllerMock) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (r *UserControllerMock) UpdateUser(ctx context.Context, id int, user domain.User) error {
	return nil
}

func (r *UserControllerMock) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (r *UserControllerMock) CheckUserEmail(ctx context.Context, user domain.User, id int) error {
	return nil
}
