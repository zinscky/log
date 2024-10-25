package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	Debug = 0
	Info  = 1
	Warn  = 2
	Error = 3
)

type Logger struct {
	LogStr []string
	Level  int
	App    string
}

func New(level int, app string) *Logger {
	return &Logger{Level: level, App: app, LogStr: []string{}}
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) Debug(format string, args ...any) {
	if l.Level <= 0 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.LogStr = append(l.LogStr, fmt.Sprintf("%s %s [%s] %s", now, "DEBUG", l.App, msg))
	}
}

func (l *Logger) Info(format string, args ...any) {
	if l.Level <= 1 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.LogStr = append(l.LogStr, fmt.Sprintf("%s %s [%s] %s", now, "INFO", l.App, msg))
	}
}

func (l *Logger) Warn(format string, args ...any) {
	if l.Level <= 2 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.LogStr = append(l.LogStr, fmt.Sprintf("%s %s [%s] %s", now, "WARN", l.App, msg))
	}
}

func (l *Logger) Error(format string, args ...any) {
	if l.Level <= 3 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.LogStr = append(l.LogStr, fmt.Sprintf("%s %s [%s] %s", now, "ERROR", l.App, msg))
	}
}

func (l *Logger) String() string {
	return strings.Join(l.LogStr, "\n")
}
