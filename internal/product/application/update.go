package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductApplication) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error) {
	return service.repo.Update(ID, name, description, category, price, stock)
}
