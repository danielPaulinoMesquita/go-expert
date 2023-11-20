package main

import (
	"fmt"
	"gorm.io/gorm"
)

func alter(db gorm.DB, product Product) {
	db.First(&product, 1)
	product.Name = "New Name TV"
	db.Save(&product)
}

func deleteRow(db gorm.DB, product Product) {
	db.First(&product, 1)
	fmt.Println(product.Name)
	db.Delete(&product)
}
