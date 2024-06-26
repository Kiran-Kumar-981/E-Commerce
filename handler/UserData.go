package handler

import (
	"fmt"
	"jwt/token/creation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	CustomerId    string `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
}

func GetAll(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, []UserData{
		{
			CustomerId:    "1",
			CustomerName:  "Kirn",
			CustomerEmail: "kiran@gmail.com",
		},
		{
			CustomerId:    "2",
			CustomerName:  "singamalai",
			CustomerEmail: "king@gmail.com",
		},
	})
}

func AddUser(ctx *gin.Context) {
	var userData UserData
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		errors := []models.ErrorDetails{
			{
				ErrorType:    models.ErrorTypeValidation,
				ErrorMessage: fmt.Sprintf("%v", err),
			},
		}
		badRequest(ctx, http.StatusBadRequest, "Invalid Product Data", errors)
		return
	}
	ok(ctx, http.StatusOK, "User Added Successfully", userData)
}
