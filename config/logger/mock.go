package logger

import (
	"sync"
)

type CapturedLogs struct {
	Debugs []string
	Infos  []string
	Warns  []string
	mu     sync.Mutex
}

type MockLogger struct {
	Logs *CapturedLogs
}

func NewMockLogger() (*MockLogger, *CapturedLogs) {
	logs := &CapturedLogs{}
	return &MockLogger{Logs: logs}, logs
}

func (l *MockLogger) Debug(msg string, args ...any) {
	l.Logs.mu.Lock()
	defer l.Logs.mu.Unlock()
	l.Logs.Debugs = append(l.Logs.Debugs, msg)
}

func (l *MockLogger) Info(msg string, args ...any) {
	l.Logs.mu.Lock()
	defer l.Logs.mu.Unlock()
	l.Logs.Infos = append(l.Logs.Infos, msg)
}

func (l *MockLogger) Warn(msg string, args ...any) {
	l.Logs.mu.Lock()
	defer l.Logs.mu.Unlock()
	l.Logs.Warns = append(l.Logs.Warns, msg)
}
