package api

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/velicanercan/simple-user-mgmt/controller"
	"github.com/velicanercan/simple-user-mgmt/docs"
)

// UserRoutes is a struct that holds the Handler and Controller
type UserRoutes struct {
	Handler    GinRouter
	Controller controller.UserController
}

// NewUserRoutes returns a UserRoutes instance
func NewUserRoutes(handler GinRouter, controller controller.UserController) UserRoutes {
	return UserRoutes{
		Handler:    handler,
		Controller: controller,
	}
}

// RegisterRoutes registers the user routes
func (r *UserRoutes) RegisterRoutes() {
	r.Handler.Gin.POST("/users", r.Controller.CreateUser)
	r.Handler.Gin.GET("/users/:id", r.Controller.GetUserByID)
	r.Handler.Gin.GET("/users", r.Controller.GetAllUsers)
	r.Handler.Gin.PUT("/users/:id", r.Controller.UpdateUser)
	r.Handler.Gin.DELETE("/users/:id", r.Controller.DeleteUser)
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	r.Handler.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
