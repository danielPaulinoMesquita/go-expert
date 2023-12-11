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
	"log"
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger) // <-- this middleware applies the logs for the requests
	//r.Use(LogRequest) /  native middleware for logs
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpressIn))
	// description WithValue - jwt is the value that we want, configs.TokenAuth is the context where the jwt will be inserted

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

/*
*
Explanation of middleware
made a 'Request' -> Middleware(usa os dados, faz alguma coisa, Continua) | outro middler | outro Midler -> Handler -> 'Response'
example below:
*/
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
