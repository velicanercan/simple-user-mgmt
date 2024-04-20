package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/velicanercan/simple-user-mgmt/api"
	"github.com/velicanercan/simple-user-mgmt/controller"
	"github.com/velicanercan/simple-user-mgmt/infrastructure"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository"
	"github.com/velicanercan/simple-user-mgmt/logger"
	"github.com/velicanercan/simple-user-mgmt/service"
)

func main() {
	router := api.NewGinRouter()

	ctx := context.Background()
	userRepository := repository.InitializeDB(ctx)

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoutes := api.NewUserRoutes(router, userController)
	userRoutes.RegisterRoutes()

	server := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: router.Gin,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log("error", "Failed to start the server!", err.Error())
		}
	}()

	// Gracefully shutdown the server
	HandleShutdown(server, userRepository)
}

func init() {
	logger.InitializeLogger()
	infrastructure.LoadEnv()
}

// HandleShutdown gracefully shuts down the server and closes the database connection
// when a signal is received
func HandleShutdown(server *http.Server, userRepository repository.UserRepository) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	logger.Log("info", "Server shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log("error", "Server forced to shutdown!", err.Error())
	}

	userRepository.Close(ctx)

	logger.Log("info", "Server exited properly.")
}
