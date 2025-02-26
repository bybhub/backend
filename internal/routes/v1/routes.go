package v1

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesV1(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			handler.SendSuccess(ctx, "v1", handler.Response{Resp: "ok"})
		})
	}
}
