package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"time"
)

type PurchaseDTO struct {
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Date      time.Time `json:"date" bson:"date"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

func MapPurchaseToJSON(purchase *model.Purchase) PurchaseDTO {
	return PurchaseDTO{
		ID:        purchase.ID,
		IDProduct: purchase.IDProduct,
		IDUser:    purchase.IDCustomer,
		Date:      purchase.Date,
		Units:     purchase.Units.Amount,
		Total:     purchase.Total.Value,
	}
}
