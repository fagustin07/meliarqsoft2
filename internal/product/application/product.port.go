package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

type IProductUseCase interface {
	Create(name string, description string, price float32, stock int, idSeller int) error
	Update(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) error
	Delete(ID uuid.UUID) error
	Find(name string, category string) domain.Product
	Filter(minPrice float32, maxPrice float32) domain.Product
}
