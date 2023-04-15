package application

import (
	"meliarqsoft2/internal/domain"
)

func (service *ProductService) Find(name string, category string) ([]*domain.Product, error) {
	return service.repo.Find(name, category)
}
