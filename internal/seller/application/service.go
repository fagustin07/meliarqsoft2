package application

import (
	"meliarqsoft2/internal/domain"
)

type SellerService struct {
	repo domain.ISellerRepository
}

func NewSellerService(repo domain.ISellerRepository) *SellerService {
	return &SellerService{repo: repo}
}
