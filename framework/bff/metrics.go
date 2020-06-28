package bff

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"git.gmtshenzhen.com/geeky-medic/x-lite/framework/types"
	"git.gmtshenzhen.com/geeky-medic/x-lite/metrics/prometheus"
)

func MetricsMiddleWare() func(ctx *gin.Context) {
	if os.Getenv("X-LITE_MODE") == "test" {
		return func(ctx *gin.Context) {
			ctx.Next()
		}
	}
	var qpsMetrics = prometheus.MustCounterWithLabelNames("request_qps", "method", "host", "path", "status")
	var latencyCounterMetrics = prometheus.MustGagueWithLabelNames("request_gauge_latency", "method", "host", "path", "status")
	var latencyMetrics = prometheus.MustHistogramWithLabelNames("request_latency_duration_seconds", []float64{0.1, 0.3, 0.5, 0.8, 1, 3}, "method", "path",
		"status")
	var responseMetrics = prometheus.MustCounterWithLabelNames("response_status_code", "code", "path", "size")
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		elapse := float64(time.Since(start)) / float64(time.Millisecond)

		method := ctx.Request.Method
		host := ctx.Request.Host
		path := ctx.Request.URL.Path
		status := fmt.Sprintf("%d", ctx.Writer.Status())

		// qps
		{
			qpsMetrics.With(method, host, path, status).Inc()
		}

		// latency
		{
			latencyMetrics.With(method, path, status).Observe(time.Since(start).Seconds())
			latencyCounterMetrics.With(method, host, path, status).Set(elapse)
		}

		// response code
		{
			value, ok := ctx.Get(types.ResponseStatusCode)
			if !ok {
				return
			}
			code, ok := value.(int)
			if !ok {
				return
			}
			responseMetrics.With(fmt.Sprintf("%d", code), ctx.Request.URL.Path, fmt.Sprintf("%d", ctx.Writer.Size())).Inc()
		}
	}
}
