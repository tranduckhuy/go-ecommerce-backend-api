package initialize

func Run() {
	// Initialize the configuration
	LoadConfig()

	// Initialize the logger
	InitLogger()

	// Initialize the database
	// InitMySQL()

	// Initialize the Redis connection
	// InitRedis()

	// Initialize the routes
	r := InitRouter()

	r.Run()
}
