package extend

import (
	"bytes"
	"fmt"
	"os"
	"sort"

	"github.com/sirupsen/logrus"

	"github.com/geekymedic/x-lite/logger"
)

type SimpleFormat struct{}

func (sf *SimpleFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var labels []string
	sort.SliceStable(labels, func(i, j int) bool {
		return labels[i] < labels[j]
	})
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v %s\n", entry.Level, entry.Message)
	return buf.Bytes(), nil
}

type SimpleLog struct {
	entry *logrus.Entry
}

func NewSimpleLog() *SimpleLog {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&SimpleFormat{})
	//log.SetFormatter(&logrus.TextFormatter{
	//	QuoteEmptyFields:       true,
	//	ForceColors:            true,
	//	DisableTimestamp:       true,
	//	DisableSorting:         true,
	//	DisableLevelTruncation: true,
	//})
	return &SimpleLog{entry: log.WithContext(nil)}
}

func (log *SimpleLog) Debug(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Debug(args...)
}

func (log *SimpleLog) Debugf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Debugf(format, args...)
}

func (log *SimpleLog) Info(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Info(args...)
}

func (log *SimpleLog) Infof(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Infof(format, args...)
}

func (log *SimpleLog) Warn(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Warn(args...)
}

func (log *SimpleLog) Warnf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Warnf(format, args...)
}

func (log *SimpleLog) Error(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Error(args...)
}

func (log *SimpleLog) Errorf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Errorf(format, args...)
}

func (log *SimpleLog) Fatal(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Fatal(args...)
}

func (log *SimpleLog) Fatalf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Fatalf(format, args...)
}

func (log *SimpleLog) Panic(args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Panic(args...)
}

func (log *SimpleLog) Panicf(format string, args ...interface{}) {
	if !log.entry.Logger.IsLevelEnabled(log.entry.Level) {
		return
	}
	log.entry.Panicf(format, args...)
}

func (log *SimpleLog) With(fields ...interface{}) logger.Logger {
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
	return &SimpleLog{entry: log.entry.WithFields(filedset)}
}

func (log *SimpleLog) SetLevel(level logger.Level) bool {
	if level <= logger.DebugLevel {
		log.entry.Logger.SetLevel(level)
		return true
	}
	return false
}
