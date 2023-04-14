package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/event/action"
	"meliarqsoft2/internal/product/domain/ports"
)

type DeleteProductsBySeller struct {
	repository    ports.IProductRepository
	deleteProduct *action.DeleteProductEvent
}

func NewDeleteProductsBySellerIdCommand(prodRepository ports.IProductRepository, deleteProduct *action.DeleteProductEvent) *DeleteProductsBySeller {
	return &DeleteProductsBySeller{
		repository:    prodRepository,
		deleteProduct: deleteProduct,
	}
}

func (command DeleteProductsBySeller) Execute(sellerID uuid.UUID) error {
	products, err := command.repository.GetFrom(sellerID)
	if err != nil {
		return err
	}

	for _, product := range products {
		err := command.deleteProduct.Execute(product.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
