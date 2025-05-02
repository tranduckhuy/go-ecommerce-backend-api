//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/controllers"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/repositories"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/routers/v1/product"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/services"
)

var ProductSet = wire.NewSet(
	repositories.NewProductRepository,
	services.NewProductService,
	controllers.NewProductController,
	product.NewRouter,
)

func InitProductRouter() (*product.Router, error) {
	wire.Build(ProductSet)
	return nil, nil
}
