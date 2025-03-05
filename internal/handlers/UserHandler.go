package handler

import (
	"net/http"

	"github.com/bybhub/backend/internal/models"
	"github.com/bybhub/backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	id, err := repositories.CreateNewUser(db, "user", &user)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := models.UserResponse{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}

	SendSuccessObject(ctx, "user", userResponse)
}

func GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := repositories.FindUserByID(db, "user", id)
	if err != nil {
		SendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	SendSuccessObject(ctx, "user", user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := repositories.FindAllUsers(db, "user")
	if err != nil {
		SendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	SendSuccessObject(ctx, "users", users)
}

func UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateData bson.M
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := repositories.FindUserByID(db, "user", id)
	if err != nil {
		SendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	updatedUser, err := repositories.UpdateUserByID(db, "user", id, updateData)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessObject(ctx, "user", updatedUser)
}

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	err := repositories.DeleteUserByID(db, "user", id)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			SendError(ctx, http.StatusNotFound, err.Error())
			return
		}
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
