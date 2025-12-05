package route

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)



func ResetRoute(handler *handler.ForgetPasswordHandler, route *gin.Engine)  {
	api := route.Group("/auth/api/v1")
	{
		api.POST("/reset", handler.RequestResetPasswordEmail)
		api.POST("/reset-password", handler.ResetPassword)
	}
}