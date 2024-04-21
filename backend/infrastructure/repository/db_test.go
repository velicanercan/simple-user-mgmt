package repository

import (
	"context"
	"testing"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/mocks/repository"
)

func TestUserRepositoryMock_GetUser(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	user, err := r.GetUser(context.Background(), 1)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if user == nil {
		t.Errorf("expected a user, got nil")
	}
	if user != nil && user.ID != 1 {
		t.Errorf("expected user id 1, got %d", user.ID)
	}
}

func TestUserRepositoryMock_InsertUser(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	err := r.InsertUser(context.Background(), &domain.User{})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestUserRepositoryMock_GetAllUsers(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	users, err := r.GetAllUsers(context.Background())
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if len(users) == 0 {
		t.Errorf("expected users, got none")
	}
}

func TestUserRepositoryMock_UpdateUser(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	err := r.UpdateUser(context.Background(), 1, &domain.User{})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestUserRepositoryMock_DeleteUser(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	err := r.DeleteUser(context.Background(), 1)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestUserRepositoryMock_Close(t *testing.T) {
	r := repository.NewUserRepositoryMock()
	r.Close(context.Background())
}
