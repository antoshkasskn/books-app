package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerHealthCheck(g *gin.RouterGroup) {
	g.GET("/", health)
}

func health(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/plain")
	ctx.String(http.StatusOK, "HEALTHY")
}
