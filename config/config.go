package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
	Port        string
	Environment string
}

type DbConfig struct {
	Host     string
	Database string
	User     string
	Password string
}

type Config struct {
	Db  *DbConfig
	App *App
}

var config *Config

func InitConfig() *Config {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("getting error reading config:%s", err.Error()))
	}
	var configs Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Sprintf("getting error unmarshalling configs:%s", err.Error()))
	}
	config = &configs
	return &configs
}

func GetConfigs() *Config {
	if config != nil {
		return config
	}
	config = InitConfig()
	return config
}

func (c *Config) isProduction() bool {
	return c.App.Environment == "production"
}
