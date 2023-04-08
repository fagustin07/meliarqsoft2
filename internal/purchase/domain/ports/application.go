package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

type IPurchaseManager interface {
	Create(IDProduct uuid.UUID, IDUser uuid.UUID, units int) (*domain.Purchase, error)
	Find(start time.Time, limit time.Time) ([]*domain.Purchase, error)
}
