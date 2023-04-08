package ports

import (
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

type IPurchaseRepository interface {
	Create(purchase *domain.Purchase) error
	Find(start time.Time, limit time.Time) ([]*domain.Purchase, error)
}
