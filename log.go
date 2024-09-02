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
}

func New(level int) *Logger {
	return &Logger{level: level}
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) Debug(format string, args ...any) {
	if l.level <= 0 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		if len(l.logStr) == 0 {
			l.logStr = []string{fmt.Sprintf("%s %s %s", now, "DEBUG", msg)}
		} else {
			l.logStr = append(l.logStr, fmt.Sprintf("%s %s %s", now, "DEBUG", msg))
		}
	}
}

func (l *Logger) Info(format string, args ...any) {
	if l.level <= 1 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		if len(l.logStr) == 0 {
			l.logStr = []string{fmt.Sprintf("%s %s %s", now, "INFO", msg)}
		} else {
			l.logStr = append(l.logStr, fmt.Sprintf("%s %s %s", now, "INFO", msg))
		}
	}
}

func (l *Logger) Warn(format string, args ...any) {
	if l.level <= 2 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		if len(l.logStr) == 0 {
			l.logStr = []string{fmt.Sprintf("%s %s %s", now, "WARN", msg)}
		} else {
			l.logStr = append(l.logStr, fmt.Sprintf("%s %s %s", now, "WARN", msg))
		}
	}
}

func (l *Logger) Error(format string, args ...any) {
	if l.level <= 3 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		if len(l.logStr) == 0 {
			l.logStr = []string{fmt.Sprintf("%s %s %s", now, "ERROR", msg)}
		} else {
			l.logStr = append(l.logStr, fmt.Sprintf("%s %s %s", now, "ERROR", msg))
		}
	}
}

func (l *Logger) FormatLogs() string {
	return strings.Join(l.logStr, "\n")
}
