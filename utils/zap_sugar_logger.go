package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapSugarLogger interface {
	Infow(msg string, keyValues ...any)
	Errorw(msg string, keyValues ...any)
	Fatalw(msg string, keyValues ...any)
}

func newZapSugarLogger() *zap.SugaredLogger {
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	logger := zap.New(core)
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
