package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/query"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type RegisterProductEvent struct {
	productRepository ports.IProductRepository
	existSeller       *query.ExistSeller
}

func (event RegisterProductEvent) Execute(name string, description string, category string, price float32, stock int, idSeller uuid.UUID) (*domain.Product, error) {
	err := event.existSeller.Execute(idSeller)
	if err != nil {
		return nil, err
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	newProduct, err := domain.NewProduct(newUUID, name, description, category, price, stock, idSeller)
	if err != nil {
		return nil, err
	}

	_, err = event.productRepository.Create(newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}
