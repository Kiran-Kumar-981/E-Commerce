// configures the data and gives to the ProductsAPI for further processing
package config

import (
	"log"

	//gorm.io/driver/mysql is the the driver package that integrate mysql to gorm.io

	"gorm.io/driver/mysql"

	//gorm is a Full-Featured ORM(Object-Relational Mapper)
	//gorm.io package gives the special methods which wrapes the data in a single statement

	"gorm.io/gorm"
)

var DB *gorm.DB
	
func ConnectDataBase() {
	//mysql database url
	dsn := "root:a@11189D001#@tcp(127.0.0.1:3306)/ganesh?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = database
}

