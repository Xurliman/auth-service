package config

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App      AppSettings      `mapstructure:"app"`
	Database DatabaseSettings `mapstructure:"database"`
	JWT      JWTSettings      `mapstructure:"jwt"`
	Mail     MailSettings     `mapstructure:"mail"`
}

type AppSettings struct {
	Port  int    `mapstructure:"port"`
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
}

type DatabaseSettings struct {
	Connection      string        `mapstructure:"connection"`
	Host            string        `mapstructure:"host"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Port            int           `mapstructure:"port"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
}

type JWTSettings struct {
	Expires int    `mapstructure:"expires"`
	Secret  string `mapstructure:"secret"`
}

type MailSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
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
			log.Fatal("Error reading config file: %v", zap.Error(err))
		}

		config := &AppConfig{}
		if err := v.Unmarshal(config); err != nil {
			log.Fatal("Error unmarshalling config: %v", zap.Error(err))
		}

		instance = config
	})
	return instance
}

func GetMailSettings() *MailSettings {
	return &instance.Mail
}

func GetJWTSecret() []byte {
	return []byte(instance.JWT.Secret)
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.env", "development")
	v.SetDefault("app.debug", true)

	v.SetDefault("database.connection", "postgres")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.name", "dbname")
	v.SetDefault("database.user", "dbuser")
	v.SetDefault("database.password", "dbpassword")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.sslmode", "disable")
	v.SetDefault("database.max_idle_conns", 10)
	v.SetDefault("database.max_open_conns", 100)
	v.SetDefault("database.max_conn_lifetime", 0)
	v.SetDefault("database.max_conn_idle_time", 8)

	v.SetDefault("jwt.expires", 12)
	v.SetDefault("jwt.secret", "secret")

	v.SetDefault("mail.host", "smtp.example.com")
	v.SetDefault("mail.port", 587)
	v.SetDefault("mail.user", "user")
	v.SetDefault("mail.password", "password")
	v.SetDefault("mail.from", "user@example.com")
}
