package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain/ports"
)

type ExistSeller struct {
	repository ports.ISellerRepository
}

// Execute TODO: revisar exist en repository
func (query ExistSeller) Execute(id uuid.UUID) (bool, error) {
	return query.repository.Exist(id) != nil, nil
}
