package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/black-dev-x/pos-go-api/internal/dto"
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/black-dev-x/pos-go-api/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductDBInterface
}

func NewProductHandler(db database.ProductDBInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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
