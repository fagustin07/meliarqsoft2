package application

import "meliarqsoft2/internal/product/domain"

func (service *ProductService) Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	return service.repo.Filter(minPrice, maxPrice)
}
