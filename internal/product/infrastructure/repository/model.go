package repository

import (
	"github.com/google/uuid"
	"time"
)

type ProductModel struct {
	ID          uuid.UUID `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Category    string    `json:"category" bson:"category"`
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
