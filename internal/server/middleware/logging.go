package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(looger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clentIp := c.ClientIP()

		looger.Info("Http request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("latency", latency.String()),
			zap.String("ip", clentIp),
		)
	}
}
