package main

import (
	"database/sql"
	"github.com/daniel/grpc/internal/database"
	"github.com/daniel/grpc/internal/pb"
	"github.com/daniel/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "gorm.io/driver/sqlite"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	// this the default port that evans make request
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
