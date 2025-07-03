package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ = cfg.Build()

	defer logger.Sync()
	sugarLogger = logger.Sugar()
}

func Info(format string, details ...interface{}) {
	sugarLogger.Infof(format, details)
}

func Error(format string, details ...interface{}) {
	if details == nil {
		sugarLogger.Error(format)
	} else {
		sugarLogger.Errorf(format, details)
	}
}
