package logman

func Debug(args ...any) {
	Stdout.Debug(args...)
}

func Debugf(tpl string, args ...any) {
	Stdout.Debugf(tpl, args...)
}

func Info(args ...any) {
	Stdout.Info(args...)
}

func Infof(tpl string, args ...any) {
	Stdout.Infof(tpl, args...)
}

func Warn(args ...any) {
	Stdout.Warn(args...)
}

func Warnf(tpl string, args ...any) {
	Stdout.Warnf(tpl, args...)
}

func Error(args ...any) {
	Stdout.Error(args...)
}

func Errorf(tpl string, args ...any) {
	Stdout.Errorf(tpl, args...)
}

func Panic(args ...any) {
	Stdout.Panic(args...)
}

func Panicf(tpl string, args ...any) {
	Stdout.Panicf(tpl, args...)
}

func Fatal(args ...any) {
	Stdout.Fatal(args...)
}

func Fatalf(tpl string, args ...any) {
	Stdout.Fatalf(tpl, args...)
}
