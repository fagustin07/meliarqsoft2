package model

import (
	"github.com/google/uuid"
	"math"
)

type Product struct {
	ID          uuid.UUID `json:"id,string,omitempty" bson:"id,string,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Category    string    `json:"category" bson:"category"`
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	Price       float32 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
}

type CreateProductRequest struct {
	Name        string    `json:"name" bson:"name" `
	Description string    `json:"description" bson:"description" `
	Category    string    `json:"category" bson:"category" `
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller" `
}

type ProductID struct {
	ID uuid.UUID `json:"product_id" bson:"product_id"`
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

type IProductService interface {
	Create(createReq CreateProductRequest) (ProductID, error)
	Update(ID uuid.UUID, updateReq UpdateProductRequest) error
	Find(name string, category string) ([]Product, error)
	Filter(minPrice float32, maxPrice float32) ([]Product, error)
	Delete(ID uuid.UUID) error
}
