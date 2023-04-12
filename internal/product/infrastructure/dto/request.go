package dto

import (
	"github.com/google/uuid"
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
