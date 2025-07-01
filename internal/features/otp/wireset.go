package otp

import "github.com/google/wire"

var Wireset = wire.NewSet(
	NewOtpHandler,
	NewOtpService,
	NewOtpRepository,
)
