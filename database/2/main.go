package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var product Product
	findFirst(*db, product)

	var product2 Product
	findWithWhere(*db, product2)

	var products []Product
	findAll(*db, products)

	// TODO to migrate Product as table
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	return
	//}

	// TODO create
	//db.Create(&Product{
	//	Name:  "VideoGame",
	//	Price: 2000.00,
	//})

	//TODO create batch
	//products := []Product{
	//	{Name: "Notebook", Price: 1000.00},
	//	{Name: "Mouse", Price: 70.00},
	//	{Name: "KeyBoard", Price: 400.00},
	//}
	//
	//db.Create(&products)
}
