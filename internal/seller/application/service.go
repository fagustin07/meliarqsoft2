package application

import (
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerService struct {
	repo ports.ISellerRepository
}

func NewSellerService(repo ports.ISellerRepository) *SellerService {
	return &SellerService{repo: repo}
}
