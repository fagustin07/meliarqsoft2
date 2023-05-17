package model

import (
	"github.com/google/uuid"
	"meliarqsoft2/pkg/exceptions/model"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       *Price
	Stock       *Stock
	IDSeller    uuid.UUID
}

func (prod *Product) CanConsume(units int) bool {
	return prod.Stock.Amount >= units
}

func (prod *Product) Consume(units int) error {
	newStock, err := NewStock(prod.Stock.Amount - units)
	if err != nil {
		return model.ProductWithoutStockError{ProductName: prod.Name}
	}
	prod.Stock = newStock
	return nil
}

func (prod *Product) Restore(units int) error {
	newStock, err := NewStock(prod.Stock.Amount + units)
	if err != nil {
		return model.ProductRestoreStockError{ProductName: prod.Name, Units: units}
	}
	prod.Stock = newStock
	return nil
}

type ProductJSON struct {
	ID          uuid.UUID `json:"id,string,omitempty" bson:"id,string,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Category    string    `json:"category" bson:"category"`
	Price       float32   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	IDSeller    uuid.UUID `json:"id_seller" bson:"id_seller"`
}

//go:generate mockgen -destination=../mock/productRepository.go -package=mock -source=product.go
type IProductRepository interface {
	Create(product Product) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (Product, error)
	Delete(ID uuid.UUID) error
	Find(name string, category string) ([]Product, error)
	Filter(minPrice float32, maxPrice float32) ([]Product, error)
	FindById(ID uuid.UUID) (*Product, error)
	UpdateStock(ID uuid.UUID, stock int) error
	GetFrom(sellerId uuid.UUID) ([]Product, error)
	FindBySellerAndName(seller uuid.UUID, name string) (*Product, error)
}

type IProductService interface {
	Update(ID uuid.UUID, updateReq UpdateProductRequest) error
	Find(name string, category string) ([]ProductJSON, error)
	Filter(minPrice float32, maxPrice float32) ([]ProductJSON, error)
}

type UpdateProductRequest struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	Price       float32 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
}
