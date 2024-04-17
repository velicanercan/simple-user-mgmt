package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velicanercan/simple-user-mgmt/models"
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
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.CreateUser(user); err != nil {
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
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := uc.Service.GetUserByID(idInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// GetAllUsers returns all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
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
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = idInt
	if err := uc.Service.UpdateUser(user); err != nil {
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
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := uc.Service.DeleteUser(idInt); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
