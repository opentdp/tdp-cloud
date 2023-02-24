package logman

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Global *zap.Logger

func New() {

	lvl := viper.GetString("logger.level")
	level, err := zapcore.ParseLevel(lvl)

	if err != nil {
		level = zap.WarnLevel
	}

	encoder := getEncoder()
	writeSyncer := getWriter()
	core := zapcore.NewCore(encoder, writeSyncer, level)

	Global = zap.New(core)
	defer Global.Sync()

	// 创建通用日志接口
	origin = Global.Named("origin").Sugar()

}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	return zapcore.NewJSONEncoder(encoderConfig)

}

func getWriter() zapcore.WriteSyncer {

	tofile := viper.GetBool("logger.tofile")
	stdout := viper.GetBool("logger.stdout")

	if tofile && stdout {
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(getFileLogger()),
			zapcore.AddSync(os.Stdout),
		)
	}

	if tofile && !stdout {
		return zapcore.AddSync(getFileLogger())
	}

	return zapcore.AddSync(os.Stdout)

}

func getFileLogger() *lumberjack.Logger {

	logFile := viper.GetString("logger.dir") + "/output.log"

	return &lumberjack.Logger{
		Filename:   logFile, // 日志文件位置
		MaxSize:    100,     // 单个日志文件最大值(单位：MB)
		MaxBackups: 21,      // 保留旧文件的最大个数
		MaxAge:     7,       // 保留旧文件的最大天数
		Compress:   true,    // 是否压缩/归档旧文件
	}

}
