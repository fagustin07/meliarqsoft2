package application

import (
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductApplication) Create(name string, description string, category string, price float32, stock int, idSeller uuid.UUID) (*domain.Product, error) {
	log.Println("LLEGUE A APPLICATION")
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return &domain.Product{}, err
	}
	newProduct := domain.NewProduct(newUUID, name, description, category, price, stock, idSeller)
	_, err = service.repo.Create(newProduct)
	return newProduct, err
}
