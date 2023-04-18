package product

import (
	"context"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain"
	"time"
)

func (repo MongoRepository) Create(product domain.Product) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	product.ID = newUUID

	dbProduct := mapProductToMongoModel(product)
	_, err = repo.collection.InsertOne(context.Background(), dbProduct)
	if err != nil {
		log.Println("Error saving product", err)
		return uuid.Nil, err
	}

	return newUUID, nil
}

func mapProductToMongoModel(product domain.Product) *ProductModel {
	return &ProductModel{
		ID: product.ID, Name: product.Name, Description: product.Description, Category: product.Category,
		Price: product.Price.Value, Stock: product.Stock.Amount, IDSeller: product.IDSeller, CreatedAt: time.Now(),
	}
}

func mapProductToDomainModel(product *ProductModel) (domain.Product, error) {
	newPrice, err := domain.NewPrice(product.Price)
	if err != nil {
		return domain.Product{}, err
	}

	stock, err := domain.NewStock(product.Stock)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       newPrice,
		Stock:       stock,
		IDSeller:    product.IDSeller,
	}, nil
}
