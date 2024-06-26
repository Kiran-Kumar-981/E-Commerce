package productsapi

import (
	"Product/management/Mysql/config"
	"Product/management/Mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	var handbags []models.Product
	config.DB.Find(&handbags)
	c.JSON(http.StatusOK, handbags)
}

func GetProduct(c *gin.Context) {
	var handbag models.Product
	if err := config.DB.First(&handbag, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Handbag not found"})
		return
	}
	c.JSON(http.StatusOK, handbag)
}

func UpdateProduct(c *gin.Context) {
	var handbag models.Product
	if err := config.DB.First(&handbag, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Handbag not found"})
		return
	}

	if err := c.ShouldBindJSON(&handbag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&handbag)
	c.JSON(http.StatusOK, handbag)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusNoContent, nil)
}
