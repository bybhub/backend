package v1

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesV1(api *gin.RouterGroup) {

	v1 := api.Group("/v1")
	{
		v1.POST("/users", handler.CreateUserHandler)       // handler.CreateUserHandler
		v1.GET("/users/:id", handler.GetUserHandler)       // handler.GetUserHandler
		v1.PUT("/users/:id", handler.UpdateUserHandler)    // handler.UpdateUserHandler
		v1.DELETE("/users/:id", handler.DeleteUserHandler) // handler.DeleteUserHandler
	}
}
