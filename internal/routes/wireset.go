package routes

import "github.com/google/wire"

var Wireset = wire.NewSet(NewRouter)
