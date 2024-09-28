package logger

import (
	"context"
	"go.uber.org/zap"
	"log"
	"sync"
)

// Logger ...
type Logger interface {
	// With adds the tags with the current log context
	With(ctx context.Context) Logger

	// Debug logs the keyvals with the current log context
	Debug(ctx context.Context, message string)

	// Info logs the keyvals with the current log context
	Info(ctx context.Context, message string)

	// Warn logs the keyvals with the current log context
	Warn(ctx context.Context, message string)

	// Error logs the keyvals with the current log context
	Error(ctx context.Context, message string)
}

// NewLogger ...
// TODO(ric): cfg is deprecated, remove it once all services have migrated the config to ucmKeyStructuredLogger.
func NewLogger(cfg *zap.Config) (Logger, error) {
	var err error

	if cfg == nil {
		panic("empty logger config")
	}

	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	sugar := zapLogger.Sugar()

	return &logger{
		log: sugar,
	}, nil
}

// NewNoopLogger returns a new no-op structured logger.
func NewNoopLogger() Logger {
	return &noop{}
}

type logger struct {
	log *zap.SugaredLogger
}

var (
	initOnce sync.Once

	// Log is the default singleton logger.
	Log Logger
)

// Init inits the default singleton logger.
var Init = func(cfg *zap.Config) {
	initOnce.Do(func() {
		var err error
		Log, err = NewLogger(cfg)
		if err != nil {
			panic(err)
		}
	})
}
