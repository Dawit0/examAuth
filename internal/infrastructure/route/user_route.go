package route

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func UserRoute(handler *handler.UserHandler, route *gin.Engine) {
	api := route.Group("/auth/api/v1")
	{
		api.POST("/create", handler.CreateUser)
	}
}
