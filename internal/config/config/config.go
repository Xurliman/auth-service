package config

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App      AppSettings      `mapstructure:"app"`
	Database DatabaseSettings `mapstructure:"database"`
	JWT      JWTSettings      `mapstructure:"jwt"`
}

type AppSettings struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseSettings struct {
	Connection      string        `mapstructure:"connection"`
	Host            string        `mapstructure:"host"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Port            int           `mapstructure:"port"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
}

type JWTSettings struct {
	Expires int `mapstructure:"expires"`
}

var (
	instance *AppConfig
	once     sync.Once
)

func Setup() *AppConfig {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(constants.ConfigPath)
		v.SetConfigType("yaml")

		setDefaults(v)

		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		config := &AppConfig{}
		if err := v.Unmarshal(config); err != nil {
			log.Fatalf("Error unmarshalling config: %v", err)
		}

		instance = config
	})
	return instance
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.env", "development")

	v.SetDefault("database.connection", "postgres")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.name", "market")
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "kali")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.max_idle_conns", 10)
	v.SetDefault("database.max_open_conns", 100)
	v.SetDefault("database.max_conn_lifetime", 0)
	v.SetDefault("database.max_conn_idle_time", 8)

	v.SetDefault("jwt.expires", 12)
}
