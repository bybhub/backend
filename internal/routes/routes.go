package routes

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/url", func(ctx *gin.Context) {
			handler.SendSuccess(ctx, "test", handler.Response{Resp: "ok"})
		})
		v1.GET("/db", handler.CreateUserHandler)
	}
}
