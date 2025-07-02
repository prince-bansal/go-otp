package api_key

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewApiService,
	NewApiHandler,
	NewApiKeyRepository,
)
