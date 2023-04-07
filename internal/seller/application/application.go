package application

import (
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerManager struct {
	repo ports.ISellerRepository
}

func NewSellerManager(repo ports.ISellerRepository) *SellerManager {
	return &SellerManager{repo: repo}
}
