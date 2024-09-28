package logger

import "context"

// Debug logs the keyvals with the current log context
func (lg *logger) Debug(ctx context.Context, message string) {
	lg.log.Debug(message)
}

// Info logs the keyvals with the current log context
func (lg *logger) Info(ctx context.Context, message string) {
	lg.log.Info(message)
}

// Warn logs the keyvals with the current log context
func (lg *logger) Warn(ctx context.Context, message string) {
	lg.log.Warn(message)
}

// Error logs the keyvals with the current log context
func (lg *logger) Error(ctx context.Context, message string) {
	lg.log.Error(message)
}

// With logs the keyvals with the current log context
func (lg *logger) With(_ context.Context) Logger {
	newLogger := lg.log.With()
	return &logger{
		log: newLogger,
	}
}
