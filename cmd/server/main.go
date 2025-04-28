package main

import (
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/initialize"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/routers"
)

func main() {

	initialize.InitConfig()
	initialize.InitLogger()

	r := routers.NewRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
