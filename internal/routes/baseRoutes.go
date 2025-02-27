package routes

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func baseRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("/app/frontend/template/index.html")
	router.Static("/static", "/app/frontend/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("frontend/static/favicon.ico")
	})
	router.GET("/health", func(ctx *gin.Context) {
		handler.SendSuccess(ctx, "health", handler.Response{Resp: "ok"})
	})
}
