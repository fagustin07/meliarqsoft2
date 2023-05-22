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
	err := actionEvent.productRepository.Delete(id)

	if err != nil {
		return err
	}

	// TODO: notificarle al servicio de compras que se borro el producto. si falla, restauro el producto
	falloBorradoCompras, err := err, nil

	if falloBorradoCompras != nil { // si falla borrado de compras, rollback de producto
		_, err2 := actionEvent.productRepository.Restore([]uuid.UUID{id})
		if err2 != nil {
			// se le avisa a backoffice el error para restauración manual
			return err
		}

		return err // esto responde al servicio de usuarios el error para que haga rollback al vendedor.
	}
	return nil
}
