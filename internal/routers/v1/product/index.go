package product

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/controllers"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/middlewares"
)

type Router struct {
	controller *controllers.ProductController
}

func NewRouter(controller *controllers.ProductController) *Router {
	return &Router{controller: controller}
}

func (r *Router) RegisterRoutes(router *gin.RouterGroup) {
	productGroup := router.Group("/products")
	{
		// Public routes
		productGroup.GET("", r.controller.List)
		productGroup.GET("/:id", r.controller.GetByID)

		// Authenticated routes (no distinction between user/admin)
		authGroup := productGroup.Group("")
		authGroup.Use(middlewares.Authentication())
		{
			authGroup.GET("/hehe", r.controller.List) // test
			authGroup.POST("/:id/like", r.controller.Like)
			authGroup.POST("/:id/review", r.controller.AddReview)
		}

		// Admin-only routes
		adminGroup := productGroup.Group("")
		adminGroup.Use(middlewares.Authentication(), middlewares.RequireAdmin())
		{
			adminGroup.GET("/admin", r.controller.List) // test
			adminGroup.POST("", r.controller.Create)
			adminGroup.PUT("/:id", r.controller.Update)
			adminGroup.DELETE("/:id", r.controller.Delete)
		}
	}
}
