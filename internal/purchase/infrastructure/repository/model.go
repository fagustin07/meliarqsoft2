package repository

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

type PurchaseModel struct {
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Date      time.Time `json:"date" bson:"date"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

func MapPurchaseToMongoModel(purchase *domain.Purchase) PurchaseModel {
	return PurchaseModel{
		IDProduct: purchase.IDProduct,
		IDUser:    purchase.IDUser,
		Date:      purchase.Date,
		Units:     purchase.Units,
		Total:     purchase.Total,
	}
}

func MapToPurchaseDomain(elem *PurchaseModel) *domain.Purchase {
	return domain.NewPurchase(elem.IDProduct, elem.IDUser, elem.Date, elem.Units, elem.Total)
}
