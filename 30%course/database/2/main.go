package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product gorm brings some properties or columns to manipulate the data. It is so useful
type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	// charset set up the time default for using in the table
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//var product Product
	//findFirst(*db, product)
	//
	//var product2 Product
	//findWithWhere(*db, product2)
	//
	//var products []Product
	//findAll(*db, products)
	//
	//var products2 []Product
	//findWithLimit(*db, products2)
	//
	//var products3 []Product
	//findManyProductsWhere(*db, products3)
	//
	//var products4 []Product
	//findManyProductsLike(*db, products4)
	//
	//var p Product
	//alter(*db, p)
	//
	// todo now, after inclusion of gorm model, the deletion is a soft delete or logic delete
	var p2 Product
	deleteRow(*db, p2)

	////TODO to migrate Product as table
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	return
	//}
	//
	////TODO create
	//db.Create(&Product{
	//	Name:  "VideoGame",
	//	Price: 2000.00,
	//})
	//
	////TODO create batch
	products := []Product{
		{Name: "Notebook", Price: 1000.00},
		{Name: "Mouse", Price: 70.00},
		{Name: "KeyBoard", Price: 400.00},
	}
	//
	db.Create(&products)
}
