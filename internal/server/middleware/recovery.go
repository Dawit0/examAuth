package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func RecoveryMiddleware(logger *zap.Logger)  gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func ()  {
			if err := recover(); err != nil {
				logger.Error("Panic recovered", zap.Any("error", err))
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}()
		ctx.Next()
	}
}