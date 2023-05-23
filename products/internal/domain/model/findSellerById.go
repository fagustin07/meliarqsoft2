package model

import "github.com/google/uuid"

//go:generate mockgen -destination=../mock/findSellerService.go -package=mock -source=findSellerById.go
type IFindSellerByIdService interface {
	Execute(id uuid.UUID) (SellerDTO, error)
}

type SellerDTO struct {
	ID           uuid.UUID `json:"seller_id" bson:"seller_id"`
	BusinessName string    `json:"businessName" bson:"businessName"`
	Email        string    `json:"email" bson:"email"`
}
