package handler

import (
	"net/http"

	"github.com/bybhub/backend/internal/models"
	"github.com/bybhub/backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Estrutura para armazenar a referência ao banco
type UserHandler struct {
	DB *mongo.Database
}

// Criar um novo usuário
func (h *UserHandler) CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := repositories.CreateNewUser(h.DB, "user", &user); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "User created successfully"})
}

// Buscar usuário por ID
func (h *UserHandler) GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := repositories.FindUserByID(h.DB, "user", id)
	if err != nil {
		SendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	SendSuccess(ctx, "user", user)
}

// Atualizar usuário por ID
func (h *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateData bson.M

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := repositories.UpdateUserByID(h.DB, "user", id, updateData); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "User updated successfully"})
}

// Deletar usuário por ID
func (h *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := repositories.DeleteUserByID(h.DB, "user", id); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "User deleted successfully"})
}
