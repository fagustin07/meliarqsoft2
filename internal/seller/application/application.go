package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerService struct {
	repo           ports.ISellerRepository
	productService *application.ProductApplication
}

func (service SellerService) Exist(seller uuid.UUID) error {
	return service.repo.Exist(seller)
}

func NewSellerManager(repo ports.ISellerRepository, productService *application.ProductApplication) *SellerService {
	return &SellerService{repo: repo, productService: productService}
}
