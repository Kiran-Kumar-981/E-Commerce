package models

//gorm is a Full-Featured ORM(Object-Relational Mapper)
//gorm.io package gives the special methods which wrapes the data in a single statement

import "gorm.io/gorm"

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Brand string  `json:"brand"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func Transfer(db *gorm.DB) {
	//send's the empty interface to insert/update the products of different data types in single place or object
	db.AutoMigrate(&Product{})
}
