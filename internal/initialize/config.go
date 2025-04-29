package initialize

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// Logger configuration
	Logger LogConfig `mapstructure:"logger"`
}

type LogConfig struct {
	Level      string `mapstructure:"log_level"`
	OutputPath string `mapstructure:"output_path"`
	Encoding   string `mapstructure:"encoding"`
}

var GlobalConfig Config

func InitConfig() {
	v := viper.New()
	v.SetConfigName("local")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = v.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
}
