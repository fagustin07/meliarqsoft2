package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductApplication) Update(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) (*domain.Product, error) {
	return service.repo.Update(ID, name, description, price, stock, idSeller)
}
