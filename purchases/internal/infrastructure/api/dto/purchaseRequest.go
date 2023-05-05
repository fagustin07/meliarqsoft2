package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"time"
)

type CreatePurchaseRequest struct {
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty" binding:"required"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty" binding:"required"`
	Units     int       `json:"units" bson:"units" binding:"required,gt=0"`
	Total     float32   `json:"total" bson:"total" binding:"required,gt=0"`
}

func (dto CreatePurchaseRequest) MapToModel() (model.Purchase, error) {
	newUnits, err := model.NewUnits(dto.Units)
	if err != nil {
		return model.Purchase{}, err
	}

	newTotal, err := model.NewTotal(dto.Total)
	if err != nil {
		return model.Purchase{}, err
	}

	return model.Purchase{
		ID:        uuid.UUID{},
		IDProduct: dto.IDProduct,
		IDUser:    dto.IDUser,
		Date:      time.Now(),
		Units:     newUnits,
		Total:     newTotal,
	}, nil
}
