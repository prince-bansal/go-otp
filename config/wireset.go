package config

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewAppConfig)
