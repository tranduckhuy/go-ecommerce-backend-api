package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/controllers"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.NewPongController().PongV1)
		v1.GET("/hello/:name", controllers.NewPongController().HelloV1)
		v1.GET("/users/:id", controllers.NewUsersController().GetUserByID)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/ping", controllers.NewPongController().PongV2)
	}

	return r
}
