package logman

import (
	"go.uber.org/zap"
)

var origin *zap.SugaredLogger

func Debug(args ...any) {
	origin.Debugln(args...)
}

func Debugf(tpl string, args ...any) {
	origin.Debugf(tpl, args...)
}

func Info(args ...any) {
	origin.Infoln(args...)
}

func Infof(tpl string, args ...any) {
	origin.Infof(tpl, args...)
}

func Warn(args ...any) {
	origin.Warnln(args...)
}

func Warnf(tpl string, args ...any) {
	origin.Warnf(tpl, args...)
}

func Error(args ...any) {
	origin.Errorln(args...)
}

func Errorf(tpl string, args ...any) {
	origin.Errorf(tpl, args...)
}

func Panic(args ...any) {
	origin.Panicln(args...)
}

func Panicf(tpl string, args ...any) {
	origin.Panicf(tpl, args...)
}

func Fatal(args ...any) {
	origin.Fatalln(args...)
}

func Fatalf(tpl string, args ...any) {
	origin.Fatalf(tpl, args...)
}
