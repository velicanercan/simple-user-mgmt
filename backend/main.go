package main

import (
	"os"

	"github.com/velicanercan/simple-user-mgmt/api"
	"github.com/velicanercan/simple-user-mgmt/controller"
	"github.com/velicanercan/simple-user-mgmt/infrastructure"
	"github.com/velicanercan/simple-user-mgmt/repository"
	"github.com/velicanercan/simple-user-mgmt/service"
)

func main() {
	router := infrastructure.NewGinRouter()

	dbc := infrastructure.InitializeDB()
	defer dbc.Close()
	
	userRepository := repository.NewUserRepository(dbc)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoutes := api.NewUserRoutes(router, userController)
	userRoutes.RegisterRoutes()

	router.Gin.Run(":" + os.Getenv("SERVER_PORT"))

}

func init() {
	infrastructure.InitializeLogger()
	infrastructure.LoadEnv()
}
