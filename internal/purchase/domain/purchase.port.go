package domain

import (
	"github.com/google/uuid"
	domain2 "meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/user/domain"
)

type IPurchaseRepository interface {
	Create(IDProduct uuid.UUID, IDUser uuid.UUID) error
	Delete(IDProduct uuid.UUID, IDUser uuid.UUID) error
	Find(IDProduct uuid.UUID, IDUser uuid.UUID) (domain2.Product, domain.User)
}
