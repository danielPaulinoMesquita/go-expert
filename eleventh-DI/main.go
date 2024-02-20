package main

import (
	"database/sql"
	"fmt"
	product2 "github.com/daniel/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// create a new product repository
	repository := product2.NewProductRepository(db)

	// create a new product usecase
	usecase := product2.NewProductUseCase(repository)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
