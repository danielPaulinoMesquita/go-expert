package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	// todo Starting example of Locking pessimist
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err)
	}

	c.Name = "Daniel Eletro"
	tx.Debug().Save(&c)
	tx.Commit()

	//2023/11/23 01:04:30 /x/x/x/x/database/lock-database/main.go:42
	//[0.324ms] [rows:1] SELECT * FROM `categories` WHERE `categories`.`id` = 1 ORDER BY `categories`.`id` LIMIT 1 FOR UPDATE
	//
	//2023/11/23 01:04:30 /x/x/x/x/database/lock-database/main.go:49
	//[1.011ms] [rows:1] UPDATE `categories` SET `name`='Daniel Eletro' WHERE `id` = 1
	//
	//Process finished with the exit code 0

}
