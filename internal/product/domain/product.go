package domain

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       float32
	Stock       int
	IDSeller    int
}

func NewProduct(id uuid.UUID, name string, description string, category string, price float32, stock int, IDSeller int) *Product {
	return &Product{ID: id, Name: name, Description: description, Category: category, Price: price, Stock: stock, IDSeller: IDSeller}
}
