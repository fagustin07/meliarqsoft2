package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
)

type CreateProductRequest struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Category    string    `json:"category" bson:"category"`
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller"`
}

type CreateProductResponse struct {
	ID          uuid.UUID `json:"id,string,omitempty" bson:"id,string,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Category    string    `json:"category" bson:"category"`
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller"`
}

func MapCreatedProductProductToJSON(product *domain.Product) CreateProductResponse {
	return CreateProductResponse{product.ID, product.Name, product.Description,
		product.Category, product.Price, product.Stock, product.IDSeller}
}
