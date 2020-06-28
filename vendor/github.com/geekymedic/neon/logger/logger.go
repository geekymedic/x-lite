package logger

import (
	"github.com/geekymedic/neon/version"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	With(fields ...interface{}) Logger
	SetLevel(level Level) bool
}

var log Logger

func init() {
	baseLog := logrus.New()
	if version.PRONAME != "" {
		baseLog.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{},
		})
	}
	log = NewBaseLog(baseLog)
}

func setJson() {
	log.(*BaseLog).entry.Logger.SetFormatter(&logrus.JSONFormatter{})
}

func SetLogger(newlogger Logger) {
	log = newlogger
}

func DefLogger() Logger {
	return log
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
func Error(args ...interface{}) {
	log.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func SetLevel(level Level) bool {
	return log.SetLevel(level)
}

func With(fields ...interface{}) Logger {
	return log.With(fields...)
}
