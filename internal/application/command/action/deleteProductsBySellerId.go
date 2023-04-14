package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/event/actions"
	"meliarqsoft2/internal/product/domain/ports"
)

type DeleteProductsBySellerId struct {
	repository    ports.IProductRepository
	deleteProduct *actions.DeleteProductEvent
}

func NewDeleteProductsBySellerIdCommand(prodRepository ports.IProductRepository, deleteProduct *actions.DeleteProductEvent) *DeleteProductsBySellerId {
	return &DeleteProductsBySellerId{
		repository:    prodRepository,
		deleteProduct: deleteProduct,
	}
}

func (command DeleteProductsBySellerId) Execute(sellerID uuid.UUID) error {
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
