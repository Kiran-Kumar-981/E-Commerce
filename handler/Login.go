package handler

import (
	"fmt"
	"jwt/token/creation/models"
	"jwt/token/creation/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var loginObject *models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginObject); err != nil {
		var errors []models.ErrorDetails = make([]models.ErrorDetails, 0, 1)
		errors = append(errors, models.ErrorDetails{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(ctx, http.StatusBadRequest, "invalid Request", errors)
	}
	var claim = &models.Token{}
	claim.CustomerId = "CustomerId"
	claim.CustomerName = loginObject.UserName
	claim.CustomerEmail = loginObject.Email
	claim.Audience = ctx.Request.Header.Get("referer")

	var tokenCreationTime = time.Now()
	var expirationTime = tokenCreationTime.Add(10 * time.Minute)
	tokenString, err := token.GenerateToken(claim, expirationTime)
	if err != nil {
		badRequest(ctx, http.StatusBadRequest, "Error in generating Token", []models.ErrorDetails{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
	}
	ok(ctx, http.StatusOK, "token Created", tokenString)
}
