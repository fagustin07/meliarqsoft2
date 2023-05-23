package model

import "github.com/google/uuid"

//go:generate mockgen -destination=../mock/deletePurchaseService.go -package=mock -source=deletePurchase.go
type IDeletePurchasesByProductsService interface {
	Execute(IDs []uuid.UUID) error
}
