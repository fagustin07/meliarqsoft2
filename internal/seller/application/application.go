package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerManager struct {
	repo ports.ISellerRepository
}

func (manager SellerManager) Exist(seller uuid.UUID) error {
	return manager.repo.Exist(seller)
}

func NewSellerManager(repo ports.ISellerRepository) *SellerManager {
	return &SellerManager{repo: repo}
}
