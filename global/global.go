package global

import (
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/logger"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	MySQL  *gorm.DB
)
