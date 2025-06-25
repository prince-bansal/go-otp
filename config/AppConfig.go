package config

import "github.com/prince-bansal/go-otp/internal/routes"

type AppConfig struct {
	Router *routes.Router
}

func NewAppConfig(router *routes.Router) *AppConfig {
	return &AppConfig{
		Router: router,
	}
}
