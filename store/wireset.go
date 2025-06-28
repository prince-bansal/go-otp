package store

import (
	"github.com/google/wire"
	"github.com/prince-bansal/go-otp/store/db"
)

var Wireset = wire.NewSet(
	db.NewStore,
)
