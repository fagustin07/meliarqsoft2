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
	if err := actionEvent.productRepository.Delete(id); err != nil {
		return err
	}

	// TODO: notificarle al servicio de compras que se borro el producto.
	return nil
}
