package handler

import (
	"net/http"

	"github.com/bybhub/backend/internal/models"
	"github.com/bybhub/backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	logger.Debug("teste")
	user := models.User{
		Name:  "Teste",
		Email: "joao@joao.com",
	}

	if err := repositories.CreateNewUser(db, "user", &user); err != nil {
		logger.Errorf("user creation error: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "user", Response{Resp: "ok"})
}
