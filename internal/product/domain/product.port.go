package domain

import "github.com/google/uuid"

type IProductRepository interface {
	Create(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) error
	Update(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) error
	Delete(ID uuid.UUID) error
	Find(name string, category string) Product
	Filter(minPrice float32, maxPrice float32) Product
}
