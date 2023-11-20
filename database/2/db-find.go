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

func findWithLimit(db gorm.DB, products []Product) {
	// select all
	println("BUSCAR COM LIMITE")
	db.Limit(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	println("BUSCAR COM LIMITE e OFFSET")
	// it's like a pagination, in this case will be two in two
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}

func findManyProductsWhere(db gorm.DB, products []Product) {
	println("FIND MANY PRODUCTS WITH WHERE")

	db.Where("price > ?", 500).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}

func findManyProductsLike(db gorm.DB, products []Product) {
	println("FIND MANY PRODUCTS WITH 'LIKE'")

	db.Where("name LIKE ?", "%book").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
