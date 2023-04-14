package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain/ports"
)

type ExistSeller struct {
	repository ports.ISellerRepository
}

func (query ExistSeller) Execute(id uuid.UUID) error {
	return query.repository.Exist(id)
}
