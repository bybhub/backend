package routes

import (
	v1 "github.com/bybhub/backend/internal/routes/v1"
	"github.com/gin-gonic/gin"
)

func apiRoutes(router *gin.Engine) {
	basePath := "/api"
	apiGoup := router.Group(basePath)
	v1.RoutesV1(apiGoup)
}
