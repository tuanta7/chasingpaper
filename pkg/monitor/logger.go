package monitor

import (
	"os"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
	bridge *otelzap.Core
}

func NewLogger(level ...zapcore.Level) (*Logger, error) {
	level = append(level, zapcore.InfoLevel)

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.Encoding = "json"
	cfg.Level = zap.NewAtomicLevelAt(level[0])

	zl, err := cfg.Build(
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			if os.Getenv("LOGGER_OUTPUT") == "stdout" {
				return core
			}

			//exporter, _ := otlploggrpc.New(nil, otlploggrpc.WithInsecure())

			provider := log.NewLoggerProvider()
			return otelzap.NewCore("", otelzap.WithLoggerProvider(provider))
		}),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: zl,
	}, nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Close() error {
	return l.logger.Sync()
}
