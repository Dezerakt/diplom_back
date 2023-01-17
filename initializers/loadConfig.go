package initializers

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	TokenTTL     int64  `mapstructure:"TOKEN_TTL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
