package command

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain/ports"
)

type DeleteProductsBySellerId struct {
	repository    ports.IProductRepository
	deleteProduct DeleteProduct
}

func NewDeleteProductsBySellerId(prodRepository ports.IProductRepository, deleteProduct DeleteProduct) *DeleteProductsBySellerId {
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
