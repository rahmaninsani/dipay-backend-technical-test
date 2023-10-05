package config

import (
	"github.com/spf13/viper"
)

type constant struct {
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`
	
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	
	AccessTokenSecretKey  string `mapstructure:"ACCESS_TOKEN_SECRET_KEY"`
	AccessTokenExpiresIn  int    `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenSecretKey string `mapstructure:"REFRESH_TOKEN_SECRET_KEY"`
	RefreshTokenExpiresIn int    `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

var Constant = constant{}

func LoadConstant() error {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	
	viper.AutomaticEnv()
	
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	
	err = viper.Unmarshal(&Constant)
	if err != nil {
		return err
	}
	
	return nil
}
