package adapters

import (
	"auth-api/app/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l ZapLogger) Info(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Info(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Error(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Error(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Debug(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Debug(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Warn(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Warn(fmt.Sprint(message), fields...)
}

func NewZapLogger(cfg config.LoggerConfig) *zap.Logger {
	var level zapcore.Level

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Colorize the log level
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // Use ISO8601 time format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	level.UnmarshalText([]byte(cfg.Level))
	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.Lock(os.Stdout),
		level,
	)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func convertToZapField(data interface{}) zap.Field {
	// First, assert that the underlying type implements zapcore.Field.
	field, ok := data.(zapcore.Field)
	if ok {
		return zap.Field(field)
	}

	return zap.Any("default_key", data)
}
