package initialize

import (
	"fmt"
	"time"

	"github.com/tranduckhuy/go-ecommerce-backend-api/global"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/entities"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorMySQL(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	// Initialize the PostgreSQL connection
	config := global.Config.MySQL

	dns := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	connectionString := fmt.Sprintf(dns, config.Username, config.Password, config.Host, config.Port, config.DbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorMySQL(err, "Failed to connect to MySQL database")

	global.Logger.Info("Connected to MySQL database")
	global.MySQL = db

	// Set pooling options
	SetPool(config)

	migrateTables()
}

func SetPool(config settings.MySQLConfig) {
	sqlDB, err := global.MySQL.DB()
	if err != nil {
		global.Logger.Error("Failed to get SQL DB instance", zap.Error(err))
		return
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
}

func migrateTables() {
	// Migrate the schema
	err := global.MySQL.AutoMigrate(
		&entities.Role{},
		&entities.User{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}
