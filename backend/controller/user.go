package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velicanercan/simple-user-mgmt/logger"
	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/service"
)

type UserController struct {
	Service service.UserService
}

// NewUserController returns a UserController instance
func NewUserController(service service.UserService) UserController {
	return UserController{
		Service: service,
	}
}

// TODO: Implement the proper error messages for all the endpoints
// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log("error", "Failed to bind JSON", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.CreateUser(user); err != nil {
		logger.Log("error", "Failed to create user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

// GetUserByID returns a user by id
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Log("error", "Failed to convert ID to integer", err.Error())
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := uc.Service.GetUserByID(idInt)
	if err != nil {
		logger.Log("error", "Failed to get user by ID", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// GetAllUsers returns all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		logger.Log("error", "Failed to get all users", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// UpdateUser updates a user
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Log("error", "Failed to convert ID to integer", err.Error())
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log("error", "Failed to bind JSON", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.UpdateUser(idInt, user); err != nil {
		logger.Log("error", "Failed to update user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Log("error", "Failed to convert ID to integer", err.Error())
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := uc.Service.DeleteUser(idInt); err != nil {
		logger.Log("error", "Failed to delete user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
