package middleware

import (
	"net/http"

	"jwt/token/creation/models"
	"jwt/token/creation/token" // Assuming this is the correct path

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("apikey")
		referer := ctx.Request.Header.Get("referer")
		valid, claim := token.VarifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(ctx)
		}
		ctx.Set("CustomerId", claim.CustomerId) // Use ctx.Set instead of modifying Keys directly
		ctx.Set("CustomerName", claim.CustomerName)
		ctx.Set("CustomerEmail", claim.CustomerEmail)
	}
}

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetails{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}
func Authorization(requiredCustomerEmail string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerEmail, exists := ctx.Get("CustomerEmail") // Use ctx.Get to retrieve values
		if !exists || customerEmail != requiredCustomerEmail {
			ReturnUnauthorized(ctx)
		}
	}
}
