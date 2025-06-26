package apiKey

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewApiService,
	NewApiHandler,
)
