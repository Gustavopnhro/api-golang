package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	product, err := CreateProduct("Test Product", 100, 10, "Electronics")
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Test Product")
	assert.Equal(t, product.Price, 100)
	assert.Equal(t, product.Quantity, 10)
	assert.Equal(t, product.Category, "Electronics")
}
