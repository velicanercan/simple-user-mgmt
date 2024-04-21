package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/logger"
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

// @Summary Creates a new user
// @Description Creates a new user
// @Accept json
// @Produce json
// @Param user body domain.User true "User to create"
// @Success 201 {object} domain.User "User created"
// @Failure 400 {object} error "Invalid request"
// @Failure 500 {object} error "Internal server error"
// @Router /users [post]
// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log("error", "Failed to bind JSON", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.CreateUser(c.Request.Context(), user); err != nil {
		logger.Log("error", "Failed to create user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User "User found"
// @Failure 400 {object} error "Invalid ID"
// @Failure 500 {object} error "Internal server error"
// @Router /users/{id} [get]
// GetUserByID returns a user by id
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Log("error", "Failed to convert ID to integer", err.Error())
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := uc.Service.GetUserByID(c.Request.Context(), idInt)
	if err != nil {
		logger.Log("error", "Failed to get user by ID", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Summary Get all users
// @Description Get all users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User "Users found"
// @Failure 500 {object} error "Internal server error"
// @Router /users [get]
// GetAllUsers returns all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Service.GetAllUsers(c.Request.Context())
	if err != nil {
		logger.Log("error", "Failed to get all users", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// @Summary Update a user
// @Description Update a user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.User true "User to update"
// @Success 200 {object} domain.User "User updated"
// @Failure 400 {object} error "Invalid ID"
// @Failure 500 {object} error "Internal server error"
// @Router /users/{id} [put]
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
	if err := uc.Service.UpdateUser(c.Request.Context(), idInt, user); err != nil {
		logger.Log("error", "Failed to update user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Summary Delete a user
// @Description Delete a user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User deleted"
// @Failure 400 {object} error "Invalid ID"
// @Failure 500 {object} error "Internal server error"
// @Router /users/{id} [delete]
// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Log("error", "Failed to convert ID to integer", err.Error())
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := uc.Service.DeleteUser(c.Request.Context(), idInt); err != nil {
		logger.Log("error", "Failed to delete user", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
