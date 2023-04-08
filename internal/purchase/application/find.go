package application

import (
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

func (manager PurchaseManager) Find(start time.Time, limit time.Time) ([]*domain.Purchase, error) {
	return manager.repository.Find(start, limit)
}
