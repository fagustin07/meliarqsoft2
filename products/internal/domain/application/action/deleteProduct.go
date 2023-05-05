package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type DeleteProductEvent struct {
	productRepository model.IProductRepository
}

func NewDeleteProductEvent(prodRepo model.IProductRepository) *DeleteProductEvent {
	return &DeleteProductEvent{
		productRepository: prodRepo,
	}
}

func (actionEvent DeleteProductEvent) Execute(id uuid.UUID) error {
	if _, err := actionEvent.productRepository.FindById(id); err != nil {
		return err
	}
	if err := actionEvent.productRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
