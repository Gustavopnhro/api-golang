package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavopnhro/api-golang/configs"
	"github.com/Gustavopnhro/api-golang/internal/entity"
	"github.com/Gustavopnhro/api-golang/internal/infra/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProductInstance(db)
	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := entity.CreateProduct(product.Name, product.Price, product.Quantity, product.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
