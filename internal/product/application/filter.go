package application

import (
	"math"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductApplication) Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	if maxPrice == 0 {
		maxPrice = math.MaxFloat32
	}
	return service.repo.Filter(minPrice, maxPrice)
}
