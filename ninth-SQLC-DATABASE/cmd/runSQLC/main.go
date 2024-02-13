package main

import (
	"context"
	"database/sql"
	"github.com/daniel/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	// todo The way to create categories
	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "BackEnd description", Valid: true},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}

	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	println(category.ID, category.Name, category.Description.String)
	//}

	// todo the way to Update Category
	//err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	//	ID:          "0113b44c-4879-489d-b7d9-c53baffce540",
	//	Name:        "Backend Updated",
	//	Description: sql.NullString{String: "Backend descr Updated", Valid: true},
	//})

	err = queries.DeleteCategory(ctx, "0113b44c-4879-489d-b7d9-c53baffce540")

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
