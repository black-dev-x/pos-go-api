package entity

import (
	"errors"
	"time"

	"github.com/black-dev-x/pos-go-api/pkg/entity"
)

var ErrorNameIsRequired = errors.New("Name is required")
var ErrorPriceIsInvalid = errors.New("Price is invalid")

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price <= 0 {
		return ErrorPriceIsInvalid
	}
	return nil
}
