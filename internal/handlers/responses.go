package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data Response) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s success", op),
		"data":    data,
	})
}

func SendRedirect(ctx *gin.Context, redirectURL string) {
	ctx.Redirect(http.StatusPermanentRedirect, redirectURL)
}

type ErrorResponse struct {
	Message    string `json:"message"`
	ErrorCorde string `json:"errorCord"`
}

type CreateResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Resp string `json:"resp"`
}
