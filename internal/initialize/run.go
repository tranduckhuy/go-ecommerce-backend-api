package initialize

import "github.com/tranduckhuy/go-ecommerce-backend-api/global"

func Run() {
	// Initialize the configuration
	LoadConfig()

	// Initialize the logger
	InitLogger()

	global.Logger.Info("Logger initialized")

	// Initialize the database
	InitMySQL()

	// Initialize the Redis connection
	InitRedis()

	// Initialize the routes
	r := InitRouter()

	r.Run()
}
