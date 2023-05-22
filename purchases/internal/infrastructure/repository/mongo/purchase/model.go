package purchase

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"time"
)

type PurchaseModel struct {
	ID         uuid.UUID `json:"_id" bson:"_id,omitempty"`
	IDProduct  uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDCustomer uuid.UUID `json:"id_customer" bson:"id_customer,omitempty"`
	Date       time.Time `json:"date" bson:"date"`
	Units      int       `json:"units" bson:"units"`
	Total      float32   `json:"total" bson:"total"`
	DeletedAt  time.Time `json:"deleted_at" bson:"deleted_at"`
}

func MapPurchaseToMongoModel(purchase *model.Purchase) PurchaseModel {
	return PurchaseModel{
		ID:         purchase.ID,
		IDProduct:  purchase.IDProduct,
		IDCustomer: purchase.IDCustomer,
		Date:       purchase.Date,
		Units:      purchase.Units.Amount,
		Total:      purchase.Total.Value,
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
		ID:         elem.ID,
		IDProduct:  elem.IDProduct,
		IDCustomer: elem.IDCustomer,
		Date:       time.Now(),
		Units:      newUnits,
		Total:      newTotal,
	}, nil
}
