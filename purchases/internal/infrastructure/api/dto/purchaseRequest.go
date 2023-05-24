package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"time"
)

type CreatePurchaseRequest struct {
	IDProduct   uuid.UUID `json:"id_product" bson:"id_product,omitempty" binding:"required"`
	IDUser      uuid.UUID `json:"id_user" bson:"id_user,omitempty" binding:"required"`
	ProductName string    `json:"product_name" bson:"product_name,omitempty" binding:"required"`
	SellerName  string    `json:"seller_name" bson:"seller_name,omitempty" binding:"required"`
	SellerEmail string    `json:"seller_email" bson:"seller_email,omitempty" binding:"required"`
	BuyerName   string    `json:"buyer_name" bson:"buyer_name,omitempty" binding:"required"`
	BuyerEmail  string    `json:"buyer_email" bson:"buyer_email,omitempty" binding:"required"`
	Units       int       `json:"units" bson:"units" binding:"required"`
	Total       float32   `json:"total" bson:"total" binding:"required"`
}

func (dto CreatePurchaseRequest) MapToModel() (*model.Purchase, *model.Notification, *model.Notification, error) {
	newUnits, err := model.NewUnits(dto.Units)
	if err != nil {
		return nil, nil, nil, err
	}

	newTotal, err := model.NewTotal(dto.Total)
	if err != nil {
		return nil, nil, nil, err
	}

	purchase := model.Purchase{
		ID:         uuid.Nil,
		IDProduct:  dto.IDProduct,
		IDCustomer: dto.IDUser,
		Date:       time.Now(),
		Units:      newUnits,
		Total:      newTotal,
	}

	sellerNotification, errSeller := model.NewNotification(dto.SellerName, dto.SellerEmail, dto.ProductName, "Sell")
	if errSeller != nil {
		return nil, nil, nil, errSeller
	}
	buyerNotification, errBuyer := model.NewNotification(dto.BuyerName, dto.BuyerEmail, dto.ProductName, "Purchase")
	if errBuyer != nil {
		return nil, nil, nil, errBuyer
	}
	return &purchase, sellerNotification, buyerNotification, nil
}

func (dto CreatePurchaseRequest) MapToInternal() (*CreatePurchase, error) {
	purchase, sellerNotification, buyerNotification, err := dto.MapToModel()
	if err != nil {
		return nil, err
	}

	return &CreatePurchase{
		Purchase:           purchase,
		SellerNotification: sellerNotification,
		BuyerNotification:  buyerNotification,
	}, nil
}

type CreatePurchase struct {
	Purchase           *model.Purchase
	SellerNotification *model.Notification
	BuyerNotification  *model.Notification
}
