package application

import (
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/purchase/domain/ports"
	application2 "meliarqsoft2/internal/user/application"
)

type PurchaseService struct {
	repository     ports.IPurchaseRepository
	productService *application.ProductService
	userService    *application2.UserService
}

func NewPurchaseService(repo ports.IPurchaseRepository, productService *application.ProductService,
	userService *application2.UserService) *PurchaseService {
	return &PurchaseService{repository: repo, productService: productService, userService: userService}
}
