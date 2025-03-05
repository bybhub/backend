package v1

import (
	handler "github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutesV1(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/", handler.CreateUserHandler)      // handler.CreateUserHandler
		users.GET("/", handler.GetAllUsers)             // handler.GetUserHandler
		users.GET("/:id", handler.GetUserHandler)       // handler.GetUserHandler
		users.PUT("/:id", handler.UpdateUserHandler)    // handler.UpdateUserHandler
		users.DELETE("/:id", handler.DeleteUserHandler) // handler.DeleteUserHandler
	}
}
