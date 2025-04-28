package initialize

import "github.com/tranduckhuy/go-ecommerce-backend-api/pkg/logger"

func InitLogger() {
	cfg := GlobalConfig.Logger

	logger.NewLogger(cfg.Level, cfg.OutputPath, cfg.Encoding)
}
