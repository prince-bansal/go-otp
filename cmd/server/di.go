//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/prince-bansal/go-otp/config"
	"github.com/prince-bansal/go-otp/internal/features/health"
	"github.com/prince-bansal/go-otp/internal/features/organisation"
	"github.com/prince-bansal/go-otp/internal/routes"
)

func InitDependencies() *config.AppConfig {
	wire.Build(organisation.WireSet, routes.Wireset, config.Wireset, health.Wireset)
	return nil
}
