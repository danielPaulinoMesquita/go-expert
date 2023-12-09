package main

import (
	"encoding/json"
	"github.com/devfullcycle/dan/goexpert/configs"
	"github.com/devfullcycle/dan/goexpert/internal/dto"
	"github.com/devfullcycle/dan/goexpert/internal/entity"
	"github.com/devfullcycle/dan/goexpert/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{}) // <-- Creating Table Product and User
	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)

}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// In the context of building web applications or APIs using Go, a handler is responsible
// for receiving an HTTP request, performing the necessary actions based on the request,
// and returning an appropriate response.
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput) //<- it try to decode body to &productInput

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productInput.Name, productInput.Price) // <- try to create Product with props from productInput
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
