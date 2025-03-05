package v1

import (
	"github.com/bybhub/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RoutesV1(api *gin.RouterGroup, db *mongo.Database) {
	handler := &handlers.UserHandler{DB: db}

	v1 := api.Group("/v1")
	{
		v1.POST("/users", handler.CreateUserHandler)       // Criar usuário
		v1.GET("/users/:id", handler.GetUserHandler)       // Buscar usuário
		v1.PUT("/users/:id", handler.UpdateUserHandler)    // Atualizar usuário
		v1.DELETE("/users/:id", handler.DeleteUserHandler) // Deletar usuário
	}
}
