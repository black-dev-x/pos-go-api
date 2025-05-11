package main

import (
	"encoding/json"
	"net/http"

	"github.com/black-dev-x/pos-go-api/configs"
	"github.com/black-dev-x/pos-go-api/internal/dto"
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/black-dev-x/pos-go-api/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs.LoadConfig()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productHandler := NewProductHandler(database.NewProductDB(db))
	userHandler := NewUserHandler(database.NewUserDB(db))

	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductDBInterface
}

func NewProductHandler(db database.ProductDBInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	print("Calling CreateProduct")
	var createProductDTO dto.CreateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&createProductDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(createProductDTO.Name, createProductDTO.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.ProductDB.Create(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

type UserHandler struct {
	UserDB database.UserDBInterface
}

func NewUserHandler(db database.UserDBInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.UserDB.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
