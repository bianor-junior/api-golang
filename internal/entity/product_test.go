package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.NoError(t, err)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 100, p.Price)
	assert.NotEmpty(t, p.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrIvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}

