package domain

import "github.com/google/uuid"

type Purchase struct {
	IDProduct uuid.UUID
	IDUser    uuid.UUID
}

func NewPurchase(IDProduct uuid.UUID, IDUser uuid.UUID) *Purchase {
	return &Purchase{IDProduct: IDProduct, IDUser: IDUser}
}
