package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {

	// Database
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_DBNAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	SSLMode    string `mapstructure:"SSL_MODE"`

	// Mail
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpPort      int    `mapstructure:"SMTP_PORT"`
	SenderEmail   string `mapstructure:"SENDER_EMAIL"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`

	// app
	AppPort string `mapstructure:"APP_PORT"`

	// middlewares
	SecretJWT string `mapstructure:"SECRET_JWT"`

	// oy
	BaseUrl  string `mapstructure:"BASEURL"`
	Username string `mapstructure:"USERNAME"`
	ApiKey   string `mapstructure:"API_KEY"`

	// Cloudinary
	CloudName   string `mapstructure:"CLOUD_NAME"`
	CloudKey    string `mapstructure:"CLOUD_KEY"`
	ApiSecret   string `mapstructure:"API_SECRET"`
	CloudFolder string `mapstructure:"CLOUD_FOLDER"`

	// IAK
	UsernameIak        string `mapstructure:"USERNAME_IAK"`
	ApiKeyIak          string `mapstructure:"API_KEY_IAK"`
	BaseUrlIakPostPaid string `mapstructure:"BASEURL_IAK_Post_Paid"`
}

var (
	AppConfig Config
)

func LoadConfig() *Config {
	viper.SetConfigType("env")
	viper.SetConfigName("public")
	// viper.SetConfigName("dev")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &AppConfig
}
