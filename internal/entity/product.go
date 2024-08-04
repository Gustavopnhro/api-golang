package entity

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidId              = errors.New("invalid identifier")
	ErrNameIsRequired         = errors.New("name is required")
	ErrPriceMustBePositive    = errors.New("price must be positive")
	ErrQuantityMustBePositive = errors.New("quantity must be positive when")
)

type Product struct {
	ID       ID     `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
	gorm.Model
}

func CreateProduct(name string, price int, quantity int, category string) (*Product, error) {
	product := Product{
		ID:       NewID(),
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}

	if err := product.ValidateProduct(); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) ValidateProduct() error {

	if p.ID.String() == "" {
		return ErrInvalidId
	}

	if _, err := ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price <= 0 {
		return ErrPriceMustBePositive
	}

	if p.Quantity <= 0 {
		return ErrQuantityMustBePositive
	}
	return nil
}
