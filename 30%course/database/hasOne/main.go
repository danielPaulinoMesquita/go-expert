package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

// Product gorm brings some properties or columns to manipulate the data. It is so useful
type Product struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Price    float64
	Category Category `gorm:"foreignKey:ID"`
	gorm.Model
}

func main() {
	// charset set up the time default for using in the table
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Now, it's necessary to pass AutoMigrate
	err = db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	//// TODO create Category
	//category := Category{Name: "Eletronicos"}
	//db.Create(&category)

	////TODO create Product
	//db.Create(&Product{
	//	Name:  "VideoGame",
	//	Price: 2000.00,
	//})
	//
	////TODO create Product with category
	product := Product{
		Name:     "Alicate",
		Price:    200.00,
		Category: Category{Name: "Ferramenta"},
	}
	db.Create(&product)

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Price, product.Category.Name)
	}

}
