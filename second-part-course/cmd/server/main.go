package main

import (
	"github.com/devfullcycle/dan/goexpert/configs"
	"github.com/devfullcycle/dan/goexpert/internal/entity"
	"github.com/devfullcycle/dan/goexpert/internal/infra/database"
	"github.com/devfullcycle/dan/goexpert/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{}) // <-- Creating Table Product and User

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpressIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger) // <-- this middleware applies the logs for the requests

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth)) // <-- Adding a middleware for verifying the token
		r.Use(jwtauth.Authenticator)               // <-- authenticating the token from context
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	// This can be considered a mux from Go
	// Multiplexer or router used for handling HTTP requests in a web application.
	// It is short for "HTTP request multiplexer."
	// http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", r)

}
