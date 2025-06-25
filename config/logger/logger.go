package logger

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
}
type NoopLogger struct{}

func (n NoopLogger) Debug(msg string, args ...any) {}
func (n NoopLogger) Info(msg string, args ...any)  {}
func (n NoopLogger) Warn(msg string, args ...any)  {}
