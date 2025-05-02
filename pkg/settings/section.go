package settings

type Config struct {
	// Logger configuration
	Logger LogConfig `mapstructure:"logger"`
	// PostgreSQL database configuration
	MySQL MySQLConfig `mapstructure:"mysql"`
	// Redis database configuration
	Redis RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"log_level"`
	OutputPath string `mapstructure:"output_path"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	Compress   bool   `mapstructure:"compress"`
	Encoding   string `mapstructure:"encoding"`
}

type MySQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}
