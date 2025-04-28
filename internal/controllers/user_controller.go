package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/services"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/responses"
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
		responses.Respond(c, responses.UserNotFound, nil)
		return
	}

	responses.Respond(c, responses.Success, user)
}
