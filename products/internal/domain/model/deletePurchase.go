package model

import "github.com/google/uuid"

type IDeletePurchasesByProductsService interface {
	Execute(IDs []uuid.UUID) error
}
