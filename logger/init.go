package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sugar is a zap sugared logger
var (
	Sugar        *zap.SugaredLogger
	loggerConfig zap.Config
)

func init() {
	if Sugar == nil {
		InitLogger("error")
	}
}

// InitLogger : constructor og logger
func InitLogger(lvString string) {
	var (
		err error
		lvl zapcore.Level
	)

	if lvl, err = getLoggerLevel(lvString); err != nil {
		log.Fatalln("failed to initalize logger due to:", err)
	}

	if lvl == zapcore.DebugLevel {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	loggerConfig.Level = zap.NewAtomicLevelAt(lvl)
	zapLogger, err := loggerConfig.Build(
		zap.AddStacktrace(zapcore.PanicLevel),
	)

	if err != nil {
		log.Fatalln("failed to initialize logger due to:", err)
	}

	Sugar = zapLogger.Sugar()
}

func getLoggerLevel(lvString string) (zapcore.Level, error) {
	var (
		lvl zapcore.Level
	)

	if err := lvl.UnmarshalText([]byte(lvString)); err != nil {
		return lvl, err
	}

	return lvl, nil
}
