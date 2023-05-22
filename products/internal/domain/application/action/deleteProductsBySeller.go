package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type DeleteProductsBySeller struct {
	productRepository model.IProductRepository
}

func NewDeleteProductsBySellerEvent(productRepo model.IProductRepository) *DeleteProductsBySeller {
	return &DeleteProductsBySeller{
		productRepository: productRepo,
	}
}

func (actionEvent DeleteProductsBySeller) Execute(sellerID uuid.UUID) ([]uuid.UUID, error) {
	deletedIds, err := actionEvent.productRepository.DeleteBySeller(sellerID)
	if err != nil {
		return nil, err
	}

	// TODO: notificar al servicio de compras los productos borrados
	return deletedIds, nil
}
