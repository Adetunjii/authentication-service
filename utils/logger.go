package utils

type AppLogger interface {
	Info(msg string, keyValues ...any)
	Error(msg string, err error)
	Fatal(msg string, err error)
}

type Logger struct {
	instance zapSugarLogger
}

func NewLogger() *Logger {
	return &Logger{
		instance: newZapSugarLogger(),
	}
}

func (l Logger) Info(msg string, keyValues ...any) {
	l.instance.Infow(msg, keyValues...)
}

func (l Logger) Error(msg string, err error) {
	if err != nil {
		l.instance.Errorw(msg, "error", err.Error())
	} else {
		l.instance.Errorw(msg, "error", "unknown")
	}
}

func (l Logger) Fatal(msg string, err error) {
	if err != nil {
		l.instance.Fatalw(msg, "error", err)
	} else {
		l.instance.Fatalw(msg, "error", "unknown")
	}
}
