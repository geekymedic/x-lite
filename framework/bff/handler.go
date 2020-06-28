package bff

import "github.com/gin-gonic/gin"

type BbfHandler func(state *State)

func HttpHandler(handler BbfHandler) gin.HandlerFunc {
	return HttpHandler(handler)
}
