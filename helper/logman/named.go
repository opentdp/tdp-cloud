package logman

import (
	"os"

	"golang.org/x/exp/slog"
)

type Logger struct {
	Name string
}

func Named(name string) *Logger {
	return &Logger{Name: name}
}

func (l *Logger) Debug(msg string, args ...any) {
	args = append([]any{slog.String("Logger", l.Name)}, args...)
	slog.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	args = append([]any{slog.String("Logger", l.Name)}, args...)
	slog.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	args = append([]any{slog.String("Logger", l.Name)}, args...)
	slog.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	args = append([]any{slog.String("Logger", l.Name)}, args...)
	slog.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	args = append([]any{slog.String("Logger", l.Name)}, args...)
	slog.Error(msg, args...)
	os.Exit(1)
}
