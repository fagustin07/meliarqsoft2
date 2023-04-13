package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/purchase/domain/ports"
	application2 "meliarqsoft2/internal/user/application"
)

type PurchaseService struct {
	repository     ports.IPurchaseRepository
	productService *application.ProductService
	userService    *application2.UserService
}

func (service PurchaseService) DeleteMany(productId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPurchaseService(repo ports.IPurchaseRepository, productService *application.ProductService,
	userService *application2.UserService) *PurchaseService {
	return &PurchaseService{repository: repo, productService: productService, userService: userService}
}
