package database

import (
	"fmt"
	"testing"

	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, error := entity.NewProduct("Test Product", 10.0)
	assert.Nil(t, error)

	error = productDB.Create(product)
	assert.Nil(t, error)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	for i := 0; i < 100; i++ {
		product, error := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, error)
		productDB.Create(product)
	}

	products, error := productDB.FindAll(0, 200, "asc")
	assert.Nil(t, error)
	assert.Len(t, *products, 100)
}
