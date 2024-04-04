package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

// Product gorm brings some properties or columns to manipulate the data. It is so useful
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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

	// todo CREATE MANY CATEGORIES TO MANY PRODUCTS
	//category := Category{Name: "Aparelho"}
	//db.Create(&category)
	//
	//category2 := Category{Name: "Ferramenta"}
	//db.Create(&category2)
	//
	//db.Create(&Product{
	//	Name:       "Celular",
	//	Price:      9000.00,
	//	Categories: []Category{category, category2},
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}
	}

}
