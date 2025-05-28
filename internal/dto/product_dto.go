package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
