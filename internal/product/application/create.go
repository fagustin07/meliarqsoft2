package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductService) Create(name string, description string, category string, price float32, stock int, idSeller uuid.UUID) (*domain.Product, error) {
	err := service.sellerService.Exist(idSeller)
	if err != nil {
		return nil, err
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return &domain.Product{}, err
	}

	newProduct, err := domain.NewProduct(newUUID, name, description, category, price, stock, idSeller)
	if err != nil {
		return &domain.Product{}, err
	}

	_, err = service.repo.Create(newProduct)

	return newProduct, err
}
