package api

import (
	"fmt"
	"os"
)

var DefaultLogger = &PrettyLogger{}

type PrettyLogger struct{}

func (l *PrettyLogger) Errorf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "\033[31m[ERROR]\033[0m "+format+"\n", args...)
}

func (l *PrettyLogger) Infof(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "\033[34m[INFO]\033[0m "+format+"\n", args...)
}

func (l *PrettyLogger) Successf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "\033[32m[SUCCESS]\033[0m "+format+"\n", args...)
}

func (l *PrettyLogger) PrintIndentf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "\033[90m||===>\033[0m "+format+"\n", args...)
}
