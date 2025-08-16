package middlewares

import (
	"book-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		body, _ := ctx.GetRawData()
		logger.FromContext(ctx).WithFields(map[string]interface{}{
			"ip":         ctx.ClientIP(),
			"method":     ctx.Request.Method,
			"path":       ctx.Request.URL.Path,
			"user_agent": ctx.Request.UserAgent(),
			"uri":        ctx.Request.RequestURI,
			"body":       string(body),
		}).Infof("Handle request")
	}
}
