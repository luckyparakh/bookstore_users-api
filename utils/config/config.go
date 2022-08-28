package config

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"

	"github.com/spf13/viper"
)

type Config struct {
	DB DBConfig
}
type DBConfig struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbAddress  string
}

func LoadConfig() (*Config, *errors.RestErr) {
	var dbconfig *Config
	viper.AddConfigPath("./utils/config/") // relative to main
	viper.SetConfigName("mysql")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.NewBadRequestError("config file not found")
	}
	if err := viper.Unmarshal(&dbconfig); err != nil {
		return nil, errors.NewBadRequestError("not able to read config file")
	}
	return dbconfig, nil
}
