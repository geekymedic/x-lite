package bff

import (
	"github.com/gin-gonic/gin"

	"git.gmtshenzhen.com/geeky-medic/x-lite/framework/bff/codes"
)

var (
	_engine = gin.Default()
	_group  = _engine.Group("/api")
)

func init() {
	_group.Use(
		MetricsMiddleWare(),
		RequestTraceMiddle(map[string]interface{}{
			"Code":    codes.CodeInternalServer,
			"Message": codes.GetMessage(codes.CodeInternalServer),
		}))
}

func RouterGroup() *gin.RouterGroup {
	return _group
}

func Engine() *gin.Engine {
	return _engine
}
