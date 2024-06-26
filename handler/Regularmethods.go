package handler

import (
	"jwt/token/creation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ok(ctx *gin.Context, Status int, Message string, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, models.Response{
		Data:    data,
		Status:  Status,
		Message: Message,
	})
}
func badRequest(ctx *gin.Context, Status int, Message string, errors []models.ErrorDetails) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
		Error:   errors,
		Status:  Status,
		Message: Message,
	})
}
