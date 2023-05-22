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

	// TODO: notificar al servicio de compras los productos borrados. si falla, restaurar productos.
	falloBorradoCompras, err := err, nil

	if falloBorradoCompras != nil { // si falla borrado, comienzo rollback
		_, err := actionEvent.productRepository.Restore(deletedIds)
		if err != nil {
			// se le avisa a backoffice el error para restauración manual
			return nil, err
		}

		return nil, err // esto responde al servicio de usuarios el error para que haga rollback al vendedor.
	}

	return deletedIds, nil
}
