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
	newUnits, err := model.NewUnits(elem.Units)
	if err != nil {
		return nil, err
	}
	newTotal, err := model.NewTotal(elem.Total)
	if err != nil {
		return nil, err

	}

	return &model.Purchase{
		ID:        elem.ID,
		IDProduct: elem.IDProduct,
		IDUser:    elem.IDUser,
		Date:      time.Now(),
		Units:     newUnits,
		Total:     newTotal,
	}, nil
}
