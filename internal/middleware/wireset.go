package middleware

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewMiddleware)
