package log

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
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

type FileLogger struct {
	Level    int
	App      string
	Filename string
}

func NewLogger(ctx any, level int) *FileLogger {
	requestId := ""
	app := ""
	switch ctx := ctx.(type) {
	case *fiber.Ctx:
		requestId = ctx.Get("X-Request-Id")
		app = ctx.Get("X-App-Name")
	case *http.Request:
		requestId = ctx.Header.Get("X-Request-Id")
		app = ctx.Header.Get("X-App-Name")
	default:
		requestId = strconv.Itoa(int(time.Now().UnixMicro()))
		app = "default-app"
	}
	return &FileLogger{
		Level:    level,
		App:      app,
		Filename: fmt.Sprintf("%s-%s.log", app, requestId),
	}
}

func (l *FileLogger) SetLevel(level int) {
	l.Level = level
}

func logToFile(filename, msg string) {
	if f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		defer f.Close()
		_, err := f.WriteString(msg + "\n")
		if err != nil {
			fmt.Println("Error writing to log file:", err)
		}
	} else {
		fmt.Println("Error opening log file:", err)
	}
}

func (l *FileLogger) Debug(format string, args ...any) {
	if l.Level <= 0 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		logToFile(l.Filename, fmt.Sprintf("%s %s [%s] %s", now, "DEBUG", l.App, msg))
	}
}

func (l *FileLogger) Info(format string, args ...any) {
	if l.Level <= 1 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		logToFile(l.Filename, fmt.Sprintf("%s %s [%s] %s", now, "INFO", l.App, msg))
	}
}

func (l *FileLogger) Warn(format string, args ...any) {
	if l.Level <= 2 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		logToFile(l.Filename, fmt.Sprintf("%s %s [%s] %s", now, "WARN", l.App, msg))
	}
}

func (l *FileLogger) Error(format string, args ...any) {
	if l.Level <= 3 {
		msg := fmt.Sprintf(format, args...)
		now := time.Now().Format(time.RFC3339)
		logToFile(l.Filename, fmt.Sprintf("%s %s [%s] %s", now, "ERROR", l.App, msg))
	}
}
