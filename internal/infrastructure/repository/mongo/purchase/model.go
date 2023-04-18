package purchase

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"time"
)

type PurchaseModel struct {
	ID        uuid.UUID `json:"_id" bson:"_id,omitempty"`
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Date      time.Time `json:"date" bson:"date"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

func MapPurchaseToMongoModel(purchase *model.Purchase) PurchaseModel {
	return PurchaseModel{
		ID:        purchase.ID,
		IDProduct: purchase.IDProduct,
		IDUser:    purchase.IDUser,
		Date:      purchase.Date,
		Units:     purchase.Units.Amount,
		Total:     purchase.Total.Value,
	}
}

func MapToPurchaseDomain(elem *PurchaseModel) (*model.Purchase, error) {
	return model.NewPurchase(elem.ID, elem.IDProduct, elem.IDUser, elem.Date, elem.Units, elem.Total)
}
