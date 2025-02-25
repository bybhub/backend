package routes

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	baseRoutes(router)
	apiRoutes(router)
}
