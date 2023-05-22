package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type FindSellerByIdEvent struct {
	repository model.ISellerRepository
}

func NewFindSellerByIdEvent(repository model.ISellerRepository) *FindSellerByIdEvent {
	return &FindSellerByIdEvent{
		repository: repository,
	}
}

func (queryEvent FindSellerByIdEvent) Execute(id uuid.UUID) (*model.Seller, error) {
	return queryEvent.repository.FindById(id)
}
