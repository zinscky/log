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
	logStr []string
	level  int
	app    string
}

func New(level int, app string) *Logger {
	return &Logger{level: level, app: app, logStr: []string{}}
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) Debug(format string, args ...any) {
	if l.level <= 0 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.logStr = append(l.logStr, fmt.Sprintf("%s %s [%s] %s", now, "DEBUG", l.app, msg))
	}
}

func (l *Logger) Info(format string, args ...any) {
	if l.level <= 1 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.logStr = append(l.logStr, fmt.Sprintf("%s %s [%s] %s", now, "INFO", l.app, msg))
	}
}

func (l *Logger) Warn(format string, args ...any) {
	if l.level <= 2 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.logStr = append(l.logStr, fmt.Sprintf("%s %s [%s] %s", now, "WARN", l.app, msg))
	}
}

func (l *Logger) Error(format string, args ...any) {
	if l.level <= 3 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		l.logStr = append(l.logStr, fmt.Sprintf("%s %s [%s] %s", now, "ERROR", l.app, msg))
	}
}

func (l *Logger) String() string {
	return strings.Join(l.logStr, "\n")
}
