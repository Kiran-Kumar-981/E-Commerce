// main.go
package main

import (
	productsapi "Product/management/Mysql/ProductsAPI" //CRUD operations are Created Retrived Updated and Deleated in this package
	"Product/management/Mysql/config"                  //database connectivity for the operations in api
	"Product/management/Mysql/models"                  //

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//we calling or trrigers the database from here
	config.ConnectDataBase()
	models.Transfer(config.DB)
	// as the naming convension suggests the given below CRUD operations are done when trigered
	router.POST("/handbags", productsapi.CreateProduct)
	router.GET("/handbags", productsapi.GetProduct)
	router.GET("/handbags/:id", productsapi.GetProduct)
	router.PUT("/handbags/:id", productsapi.UpdateProduct)
	router.DELETE("/handbags/:id", productsapi.DeleteProduct)

	router.Run(":2222")
}
