package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tranduckhuy/go-ecommerce-backend-api/global"
)

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("local")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err = v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
}
