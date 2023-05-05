package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
)

type DeleteProductsBySeller struct {
	repository    model.IProductRepository
	findProducts  *query.FindProductsBySeller
	deleteProduct *DeleteProductEvent
}

func NewDeleteProductsBySellerCommand(repository model.IProductRepository,
	findProducts *query.FindProductsBySeller,
	deleteProduct *DeleteProductEvent) *DeleteProductsBySeller {

	return &DeleteProductsBySeller{
		repository:    repository,
		findProducts:  findProducts,
		deleteProduct: deleteProduct,
	}
}

func (actionCommand DeleteProductsBySeller) Execute(sellerID uuid.UUID) error {
	products, err := actionCommand.findProducts.Execute(sellerID)
	if err != nil {
		return err
	}

	for _, product := range products {
		err := actionCommand.deleteProduct.Execute(product.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
