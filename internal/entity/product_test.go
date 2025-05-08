package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Test Product", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Test Product", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorNameIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	p, err := NewProduct("Test Product", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorPriceIsInvalid, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	p, err := NewProduct("Test Product", -1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorPriceIsInvalid, err)
	assert.Nil(t, p)
}
