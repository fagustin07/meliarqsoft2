package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

type ProductService struct {
	repo domain.IProductRepository
}

func NewProductUseCase(repo domain.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (service *ProductService) Create(name string, description string, price float32, stock int, idSeller int) (*domain.Product, error) {
	newUUID, err := uuid.NewUUID()
	if err == nil {
		return &domain.Product{}, err
	}
	return service.repo.Create(newUUID, name, description, price, stock, idSeller)
}
