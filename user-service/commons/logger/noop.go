package logger

import "context"

// NoOpLogger this is a noop structured logger. Doesn't do anything.
type noop struct{}

// With ...
func (n *noop) With(ctx context.Context) Logger {
	return n
}

// Debug ...
func (*noop) Debug(ctx context.Context, message string) {
	// Do nothing
}

// Info ...
func (*noop) Info(ctx context.Context, message string) {
	// Do nothing
}

// Warn ...
func (*noop) Warn(ctx context.Context, message string) {
	// Do nothing
}

// Error ...
func (*noop) Error(ctx context.Context, message string) {
	// Do nothing
}
