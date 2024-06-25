package main

import (
	"jwt/token/creation/handler"
	"jwt/token/creation/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", handler.LoginHandler)
	router.POST("/api/product", middleware.ValidateToken(), handler.GetAll)
	http.ListenAndServe(":1111", router)
}
