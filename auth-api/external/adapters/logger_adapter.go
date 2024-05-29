package adapters

import (
	"auth-api/app/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLogger struct {
	cfg    config.LoggerConfig
	logger *zap.Logger
}

func (l ZapLogger) Info1(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Info(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Error1(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Error(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Debug1(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Debug(fmt.Sprint(message), fields...)
}

func (l ZapLogger) Warn1(message interface{}, params ...interface{}) {
	fields := make([]zap.Field, 0, len(params))
	for _, item := range params {
		fields = append(fields, convertToZapField(item))
	}

	l.logger.Warn(fmt.Sprint(message), fields...)
}

func NewZapLogger(cfg config.LoggerConfig) *ZapLogger {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse log level: %v", err)
		os.Exit(1)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	logger := zap.New(core)

	return &ZapLogger{
		cfg:    cfg,
		logger: logger,
	}
}

func convertToZapField(data interface{}) zap.Field {
	// First, assert that the underlying type implements zapcore.Field.
	field, ok := data.(zapcore.Field)
	if ok {
		return zap.Field(field)
	}

	return zap.Any("default_key", data)
}
