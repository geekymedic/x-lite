package extend

import (
	"github.com/geekymedic/x-lite/logger"
	"google.golang.org/grpc/grpclog"
)

var defLog *GrpcLog

type GrpcLog struct {
	logger.Logger
	GrpcExtendLog
}

type GrpcExtendLog interface {
	Infoln(args ...interface{})
	Debugln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Warning(args ...interface{})
	Warnln(args ...interface{})
	V(l int) bool
}

func init() {
	defLog = &GrpcLog{Logger: logger.DefLogger()}
}

func (log *GrpcLog) Infoln(args ...interface{}) {
	log.Info(args...)
}

func (log *GrpcLog) Debugln(args ...interface{}) {
	log.Debug(args...)
}

func (log *GrpcLog) Warning(args ...interface{}) {
	log.Warning(args...)
}

func (log *GrpcLog) Warningln(args ...interface{}) {
	log.Warn(args...)
}
func (log *GrpcLog) Warningf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
func (log *GrpcLog) Warnln(args ...interface{}) {
	log.Warn(args...)
}

func (log *GrpcLog) Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func (log *GrpcLog) Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

func (log *GrpcLog) V(l int) bool {
	return log.SetLevel(logger.Level(l))
}

func ReplaceGrpcLogger() {
	grpclog.SetLoggerV2(defLog)
}

func SetGrpcLog(log logger.Logger) {
	defLog = &GrpcLog{Logger: log}
}

func DefGrpLog() *GrpcLog {
	return defLog
}
