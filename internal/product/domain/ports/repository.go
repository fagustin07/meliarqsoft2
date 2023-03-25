package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

type IProductRepository interface {
	Create(product *domain.Product) (*domain.Product, error)
	Update(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) (*domain.Product, error)
	Delete(ID uuid.UUID) error
	Find(name string, category string) ([]*domain.Product, error)
	Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error)
}
