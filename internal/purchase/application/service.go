package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
	"meliarqsoft2/internal/product/application"
	application2 "meliarqsoft2/internal/user/application"
)

type PurchaseService struct {
	repository     domain.IPurchaseRepository
	productService *application.ProductService
	userService    *application2.UserService
}

func (service PurchaseService) DeleteMany(productId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPurchaseService(repo domain.IPurchaseRepository, productService *application.ProductService,
	userService *application2.UserService) *PurchaseService {
	return &PurchaseService{repository: repo, productService: productService, userService: userService}
}
