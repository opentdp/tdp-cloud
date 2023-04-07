package logman

import (
	"os"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"

	"tdp-cloud/cmd/args"
)

func New() {

	var level slog.Level
	var handler slog.Handler

	level.UnmarshalText([]byte(args.Logger.Level))

	opt := slog.HandlerOptions{
		Level: level,
	}

	switch args.Logger.Target {
	case "file":
		handler = opt.NewJSONHandler(fileWriter())
	case "stdout":
		handler = opt.NewTextHandler(os.Stdout)
	default:
		handler = opt.NewTextHandler(os.Stderr)
	}

	slog.SetDefault(slog.New(handler))

}

func fileWriter() *lumberjack.Logger {

	logFile := args.Logger.Dir + "/output.log"

	return &lumberjack.Logger{
		Filename:   logFile, // 日志文件位置
		MaxSize:    100,     // 单个日志文件最大值(单位：MB)
		MaxBackups: 21,      // 保留旧文件的最大个数
		MaxAge:     7,       // 保留旧文件的最大天数
		Compress:   true,    // 是否压缩/归档旧文件
	}

}
