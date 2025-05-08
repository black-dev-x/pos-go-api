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

func TestUpgradeProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, error := entity.NewProduct("Test Product", 10.0)
	assert.Nil(t, error)

	error = productDB.Create(product)
	assert.Nil(t, error)

	product.Name = "Updated Product"
	error = productDB.Update(product)
	assert.Nil(t, error)

	updatedProduct, error := productDB.FindByID(product.ID.String())
	assert.Nil(t, error)
	assert.Equal(t, "Updated Product", updatedProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, error := entity.NewProduct("Test Product", 10.0)
	assert.Nil(t, error)

	error = productDB.Create(product)
	assert.Nil(t, error)

	error = productDB.Delete(product.ID.String())
	assert.Nil(t, error)

	product, error = productDB.FindByID(product.ID.String())
	assert.NotNil(t, error)
	assert.Nil(t, product)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, error := entity.NewProduct("Test Product", 10.0)
	assert.Nil(t, error)

	error = productDB.Create(product)
	assert.Nil(t, error)

	foundProduct, error := productDB.FindByID(product.ID.String())
	assert.Nil(t, error)
	assert.Equal(t, product.ID.String(), foundProduct.ID.String())
}

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
