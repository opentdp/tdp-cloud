package logman

import (
	"os"

	"golang.org/x/exp/slog"
)

type Logger struct {
	*slog.Logger
	Name string
}

func NewLogger(name string) *Logger {
	return &Logger{
		Logger: slog.Default(),
		Name:   name,
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	args = append([]any{"Logger", l.Name}, args...)
	l.Logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	args = append([]any{"Logger", l.Name}, args...)
	l.Logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	args = append([]any{"Logger", l.Name}, args...)
	l.Logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	args = append([]any{"Logger", l.Name}, args...)
	l.Logger.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.Logger.Error(msg, args...)
	os.Exit(1)
}
