package main

import (
	"os"

	"github.com/velicanercan/simple-user-mgmt/api"
	"github.com/velicanercan/simple-user-mgmt/controller"
	"github.com/velicanercan/simple-user-mgmt/infrastructure"
	"github.com/velicanercan/simple-user-mgmt/infrastructure/repository"
	"github.com/velicanercan/simple-user-mgmt/logger"
	"github.com/velicanercan/simple-user-mgmt/service"
)

func main() {
	router := infrastructure.NewGinRouter()

	dbc := repository.InitializeDB()
	defer dbc.Close()

	userService := service.NewUserService(dbc)
	userController := controller.NewUserController(userService)
	userRoutes := api.NewUserRoutes(router, userController)
	userRoutes.RegisterRoutes()

	router.Gin.Run(":" + os.Getenv("SERVER_PORT"))

}

func init() {
	logger.InitializeLogger()
	infrastructure.LoadEnv()
}
