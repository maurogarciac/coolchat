package common

import (
	"context"
	"os"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) *zap.SugaredLogger {
	writer := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	encoder.AddString("app", "backend")

	core := zapcore.NewCore(encoder, writer, getLoggingLevelBy(level))

	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

func getLoggingLevelBy(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "fatal":
		return zapcore.FatalLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn":
		return zapcore.WarnLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func GetReqID(ctx context.Context) string {
	return middleware.GetReqID(ctx)
}
