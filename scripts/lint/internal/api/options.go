package api

type Logger interface {
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
	Successf(format string, args ...any)
	PrintIndentf(format string, args ...any)
}

type Options struct {
	Logger Logger
}

type CommandOption func(*Options)

func WithLogger(logger Logger) CommandOption {
	return func(o *Options) {
		o.Logger = logger
	}
}
