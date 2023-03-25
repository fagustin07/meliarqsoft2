package application

import (
	"meliarqsoft2/internal/product/domain/ports"
)

type ProductApplication struct {
	repo ports.IProductRepository
}

func NewProductApplication(repo ports.IProductRepository) *ProductApplication {
	return &ProductApplication{repo: repo}
}
