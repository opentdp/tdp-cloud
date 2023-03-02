package logman

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"tdp-cloud/cmd/args"
)

var global *zap.Logger

func New() {

	level, err := zapcore.ParseLevel(args.Logger.Level)

	if err != nil {
		level = zap.WarnLevel
	}

	core := zapcore.NewCore(getEncoder(), getWriter(), level)

	// 创建全局接口
	global = zap.New(core)

	// 创建通用接口
	origin = Named("origin").Sugar()

}

func Sync() error {
	return global.Sync()
}

func Named(n string) *zap.Logger {
	return global.Named(n)
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	return zapcore.NewJSONEncoder(encoderConfig)

}

func getWriter() zapcore.WriteSyncer {

	if args.Logger.ToFile {
		fw := zapcore.AddSync(fileWriter())
		if args.Logger.Stdout {
			return zapcore.NewMultiWriteSyncer(
				zapcore.AddSync(os.Stdout), fw,
			)
		}
		return fw
	}

	return zapcore.AddSync(os.Stdout)

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
