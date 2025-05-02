package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/wire"
)

func InitRouter(r *gin.Engine) {
	productRouter, err := wire.InitProductRouter()
	if err != nil {
		panic("Failed to initialize product router: " + err.Error())
	}

	fmt.Println("Product router initialized successfully")
	fmt.Println("Product router:", productRouter)
	apiV1 := r.Group("/api/v1")
	{
		productRouter.RegisterRoutes(apiV1)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
