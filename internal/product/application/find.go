package application

import "meliarqsoft2/internal/product/domain"

func (service *ProductApplication) Find(name string, category string) ([]*domain.Product, error) {
	return service.repo.Find(name, category)
}
