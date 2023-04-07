package logman

import (
	"time"

	"golang.org/x/exp/slog"
)

func Any(key string, val any) slog.Attr {
	return slog.Any(key, val)
}

func Bool(key string, val bool) slog.Attr {
	return slog.Bool(key, val)
}

func Duration(key string, val time.Duration) slog.Attr {
	return slog.Duration(key, val)
}

func Float64(key string, val float64) slog.Attr {
	return slog.Float64(key, val)
}

func Group(key string, val ...slog.Attr) slog.Attr {
	return slog.Group(key, val...)
}

func Int(key string, val int) slog.Attr {
	return slog.Int(key, val)
}

func Int64(key string, val int64) slog.Attr {
	return slog.Int64(key, val)
}

func String(key string, val string) slog.Attr {
	return slog.String(key, val)
}

func Time(key string, val time.Time) slog.Attr {
	return slog.Time(key, val)
}

func Uint64(key string, val uint64) slog.Attr {
	return slog.Uint64(key, val)
}
