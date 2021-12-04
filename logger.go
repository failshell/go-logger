package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	New()
}

func New() (*zap.Logger, error) {
	level := zapcore.InfoLevel

	val, exists := os.LookupEnv("DEBUG")
	if exists && val == "1" {
		level = zapcore.DebugLevel
	}

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			ConsoleSeparator: " ",
		},
	}
	logger, _ = cfg.Build()
	defer logger.Sync()

	return logger, nil
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Fatal(msg string) {
	logger.Fatal(msg)
	os.Exit(1)
}

func Info(msg string) {
	logger.Info(msg)
}

func Warning(msg string) {
	logger.Warn(msg)
}
