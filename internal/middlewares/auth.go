package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/responses"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			responses.Respond(c, responses.Unauthorized, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
