package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/global"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/controllers"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/routers"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/wire"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	routers.InitRouter(r)

	userController, _ := wire.InitUserRouterHandler()

	v1 := r.Group("/api/test/v1")
	{
		v1.GET("/ping", controllers.NewPongController().PongV1)
		v1.GET("/hello/:name", controllers.NewPongController().HelloV1)
		v1.GET("/users/:id", userController.GetUserByID)
	}

	v2 := r.Group("/api/test/v2")
	{
		v2.GET("/ping", controllers.NewPongController().PongV2)
	}

	return r
}
