package main_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	main "github.com/velicanercan/simple-user-mgmt"
	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/logger"
)

func TestHandleShutdown(t *testing.T) {
	quit := make(chan os.Signal, 1)
	server := &http.Server{}
	userRepository := &mockUserRepository{}

	logger.InitializeLogger()
	go func() {
		time.Sleep(100 * time.Millisecond)
		quit <- os.Interrupt
	}()

	main.HandleShutdown(quit, server, userRepository)

	assert.True(t, userRepository.closed, "User repository should be closed")
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type mockUserRepository struct {
	closed bool
}

func (m *mockUserRepository) Close(ctx context.Context) {
	m.closed = true
}

func (m *mockUserRepository) InsertUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (m *mockUserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

func (m *mockUserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (m *mockUserRepository) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	return nil
}

func (m *mockUserRepository) DeleteUser(ctx context.Context, id int) error {
	return nil
}
