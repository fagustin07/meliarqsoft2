package dto

import (
	"github.com/google/uuid"
	"math"
	"meliarqsoft2/internal/domain/model"
)

type CreateProductRequest struct {
	Name        string    `json:"name" bson:"name" binding:"required"`
	Description string    `json:"description" bson:"description" binding:"required"`
	Category    string    `json:"category" bson:"category" binding:"required"`
	Price       float32   `json:"price" bson:"price" binding:"required"`
	Stock       int       `json:"stock" bson:"stock" binding:"required"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	Price       float32 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
}

func (dto CreateProductRequest) MapToModel() (model.Product, error) {
	newPrice, err := model.NewPrice(dto.Price)
	if err != nil {
		return model.Product{}, err
	}

	stock, err := model.NewStock(dto.Stock)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       newPrice,
		Stock:       stock,
		IDSeller:    dto.IDSeller,
	}, nil
}

type FilterProductQuery struct {
	MinPrice *float32 `form:"min_price"`
	MaxPrice *float32 `form:"max_price"`
}

func (filterProdQuery FilterProductQuery) GetMin() float32 {
	if filterProdQuery.MinPrice != nil {
		return *filterProdQuery.MinPrice
	} else {
		return float32(0)
	}
}

func (filterProdQuery FilterProductQuery) GetMax() float32 {
	if filterProdQuery.MaxPrice != nil {
		return *filterProdQuery.MaxPrice
	} else {
		return math.MaxFloat32
	}
}
