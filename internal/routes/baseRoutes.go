package routes

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func baseRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		handler.SendSuccess(ctx, "index", handler.Response{Resp: "ok"})
	})
	router.GET("/health", func(ctx *gin.Context) {
		handler.SendSuccess(ctx, "health", handler.Response{Resp: "ok"})
	})
}
