package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

type IProductService interface {
	Create(name string, description string, category string, price float32, stock int, idSeller uuid.UUID) (*domain.Product, error)
	Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error)
	Delete(ID uuid.UUID) error
	Find(name string, category string) ([]*domain.Product, error)
	Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error)
	ManipulateStock(idProduct uuid.UUID, units int) (float32, error)
	DeleteMany(sellerId uuid.UUID) error
}
