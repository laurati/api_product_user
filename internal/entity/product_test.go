package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	p1, err := NewProduct("p1", 10.4)
	assert.NotNil(t, p1)
	assert.Nil(t, err)
	assert.NotEmpty(t, p1.ID)
	assert.Equal(t, "p1", p1.Name)
	assert.Equal(t, 10.4, p1.Price)
}

func TestProductName(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Error(t, err, ErrNameIsRequired)
}

func TestProductPriceRequired(t *testing.T) {
	p, err := NewProduct("p", 0)
	assert.Nil(t, p)
	assert.Error(t, err, ErrPriceIsRequired)
}

func TestProductPriceInvalid(t *testing.T) {
	p, err := NewProduct("p", -1)
	assert.Nil(t, p)
	assert.Error(t, err, ErrInvalidPrice)
}

func TestProductvalidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
