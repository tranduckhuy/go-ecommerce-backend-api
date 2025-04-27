package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/services"
)

type UsersController struct {
	userService *services.UserService
}

func NewUsersController() *UsersController {
	return &UsersController{
		userService: services.NewUserService(),
	}
}

// Get user by user ID
func (u *UsersController) GetUserByID(c *gin.Context) {
	// Implementation for getting user by ID
	userID := c.Param("id")
	user, err := u.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"user": user})
}
