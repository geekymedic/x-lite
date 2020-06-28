package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type Level = logrus.Level

type BaseLog struct {
	entry *logrus.Entry
}

func NewBaseLog(log *logrus.Logger) *BaseLog {
	return &BaseLog{entry: logrus.NewEntry(log)}
}

func (log *BaseLog) Debug(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Debug(args...)
}

func (log *BaseLog) Debugf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Debugf(format, args...)
}

func (log *BaseLog) Info(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Info(args...)
}

func (log *BaseLog) Infof(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Infof(format, args...)
}

func (log *BaseLog) Warn(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Warn(args...)
}

func (log *BaseLog) Warnf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Warnf(format, args...)
}

func (log *BaseLog) Error(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Error(args...)
}

func (log *BaseLog) Errorf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Errorf(format, args...)
}

func (log *BaseLog) Fatal(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Fatal(args...)
}

func (log *BaseLog) Fatalf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Fatalf(format, args...)
}

func (log *BaseLog) Panic(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Panic(args...)
}

func (log *BaseLog) Panicf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Panicf(format, args...)
}

func (log *BaseLog) With(fields ...interface{}) Logger {
	var filedset = logrus.Fields{}
	i := 0
	for ; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			filedset[fmt.Sprintf("%s", fields[i])] = fields[i+1]
		}
	}
	if i+1 == len(fields) {
		filedset[fmt.Sprintf("%s", fields[i])] = ""
	}
	return &BaseLog{entry: log.entry.WithFields(filedset)}
}

func (log *BaseLog) SetLevel(level Level) bool {
	if level <= DebugLevel {
		log.entry.Logger.SetLevel(level)
		return true
	}
	return false
}
