package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
}

func (p *Product) ValidateProduct() error {
	if p.ID == uuid.Nil {
		return errors.New("invalid ID")
	}

	if p.Name == "" {
		return errors.New("name cannot be empty")
	}

	if p.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if _, err := time.Parse(time.RFC3339, p.CreatedAt); err != nil {
		return errors.New("invalid created_at format, must be RFC3339")
	}

	return nil
}
