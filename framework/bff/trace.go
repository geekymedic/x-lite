package bff

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/geekymedic/x-lite/framework/bff/codes"
	"github.com/geekymedic/x-lite/framework/types"
	"github.com/geekymedic/x-lite/version"
)

func RequestTraceMiddle(failOut map[string]interface{}, ignore ...string) gin.HandlerFunc {
	var skips = map[string]struct{}{}
	for _, _ignore := range ignore {
		skips[_ignore] = struct{}{}
	}
	return func(c *gin.Context) {
		state := NewState(c)
		log := state.Logger
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		defer func() {
			err := recover()
			if err != nil {
				if !c.Writer.Written() {
					c.JSON(http.StatusOK, failOut)
					c.Set(types.ResponseStatusCode, failOut["Code"])
				}
				var buf = debug.Stack()
				log.With("stack", err, "full-stack", fmt.Sprintf("%s", buf)).Error("panic stack")
				c.Abort()
				return
			}
			session := types.NewSessionFromGinCtx(c)

			// Log only when path is not being skipped
			if _, ok := skips[path]; !ok {
				param := gin.LogFormatterParams{
					Request: c.Request,
					Keys:    c.Keys,
				}

				// Stop timer
				param.TimeStamp = time.Now()
				param.Latency = param.TimeStamp.Sub(start)
				param.ClientIP = c.ClientIP()
				param.Method = c.Request.Method
				param.StatusCode = c.Writer.Status()
				param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

				param.BodySize = c.Writer.Size()

				if raw != "" {
					path = path + "?" + raw
				}
				param.Path = path

				var body []byte
				var err error
				if cb, ok := c.Get(gin.BodyBytesKey); ok {
					if cbb, ok := cb.([]byte); ok {
						body = cbb
					}
				}
				if body == nil {
					body, err = ioutil.ReadAll(c.Request.Body)
					if err != nil {
						// TODO return it
					} else {
						c.Set(gin.BodyBytesKey, body)
					}
				}

				contentSize := c.Request.Header.Get("Content-Length")
				sessionLog := session.ShortLog()
				log = log.With("pro_name", version.PRONAME, "git_commit", version.ShortGitCommit(),
					"method", param.Method,
					"status", param.StatusCode,
					"content_length", contentSize,
					"real_req_size", fmt.Sprintf("%v", len(body)),
					"resp_size", param.BodySize,
					"latency", fmt.Sprintf("%v", param.Latency),
					"client_ip", param.ClientIP,
				).With(sessionLog...)

				code, ok := c.Get(types.ResponseStatusCode)
				if ok {
					log = log.With("err_code", code)
				}
				msg, ok := c.Get(types.ResponseErr)
				if ok {
					log = log.With("trace_msg", msg)
				}
				log = log.With("inbound", string(body))
				outbound, ok := c.Get(types.ResponseBody)
				if ok {
					log = log.With("outbound", outbound)
				}
				if code == 0 {
					log.Info("http trace log")
				} else if code == codes.CodeInternalServer {
					log.Error("http trace log")
				} else {
					log.Warn("http trace log")
				}
			}
		}()

		// Process request
		c.Next()
	}
}
