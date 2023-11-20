package main

import (
	"fmt"
	"gorm.io/gorm"
)

func findFirst(db gorm.DB, product Product) {
	//select one
	println("BUSCAR UM REGISTRO COM PARAMETRO")
	db.First(&product, 2)
	fmt.Println(product)
}

func findWithWhere(db gorm.DB, product Product) {
	// select with "where"
	println("BUSCAR UM REGISTRO COM PARAMETRO")
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)

}

func findAll(db gorm.DB, products []Product) {
	// select all
	println("BUSCAR TODAS")
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
