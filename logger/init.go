package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sugar is a zap sugared logger
var Sugar *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction(
		zap.AddStacktrace(zapcore.PanicLevel),
	)
	if err != nil {
		log.Fatalln("failed to initialize zap logger")
	}
	Sugar = logger.Sugar()
}

// AddDebugLogger is to add new logger for debug
func AddDebugLogger() {
	l, err := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.PanicLevel),
	)
	if err != nil {
		Sugar.Fatal("failed to initialize zap logger")
	}
	Sugar = l.Sugar()
	Sugar.Infof("Debug is on\n")
}
