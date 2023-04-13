package command

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain/ports"
)

type DeleteSeller struct {
	sellerRepository         ports.ISellerRepository
	deleteProduct            DeleteProduct
	deleteProductsBySellerId DeleteProductsBySellerId
}

func NewDeleteSeller(sellerRepo ports.ISellerRepository, deleteProd DeleteProduct) *DeleteSeller {
	return &DeleteSeller{
		sellerRepository: sellerRepo,
		deleteProduct:    deleteProd,
	}
}

func (usecase DeleteSeller) Execute(id uuid.UUID) error {
	if err := usecase.sellerRepository.Exist(id); err != nil {
		return err
	}

	if err := usecase.deleteProductsBySellerId.Execute(id); err != nil {
		return err
	}

	if err := usecase.sellerRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
