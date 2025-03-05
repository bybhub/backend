package handler

import (
	"net/http"

	"github.com/bybhub/backend/internal/models"
	"github.com/bybhub/backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Criar um novo usuário
func CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := repositories.CreateNewUser(db, "user", &user); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "User created successfully"})
}

// Buscar usuário por ID
func GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := repositories.FindUserByID(db, "user", id)
	if err != nil {
		SendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: user.Email})
}

// Atualizar usuário por ID
func UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateData bson.M

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := repositories.UpdateUserByID(db, "user", id, updateData); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "User updated successfully"})
}

// Deletar usuário por ID
func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := repositories.DeleteUserByID(db, "user", id); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "ok"})
}
