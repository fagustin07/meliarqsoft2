package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RestoreProductsEvent struct {
	productRepository model.IProductRepository
}

func NewRestoreProductsEvent(productRepository model.IProductRepository) *RestoreProductsEvent {
	return &RestoreProductsEvent{
		productRepository: productRepository,
	}
}

func (actionEvent RestoreProductsEvent) Execute(IDs []uuid.UUID) (int64, error) {
	countRestored, err := actionEvent.productRepository.Restore(IDs)
	if err != nil {
		return 0, err
	}
	return countRestored, nil
}
