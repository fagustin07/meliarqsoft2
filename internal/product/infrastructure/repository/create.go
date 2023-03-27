package repository

import (
	"context"
	"log"
	"meliarqsoft2/internal/product/domain"
	"time"
)

func (repo ProductMongoDBRepository) Create(product *domain.Product) (*domain.Product, error) {
	dbProduct := mapProductToMongoModel(product)
	_, err := repo.collection.InsertOne(context.Background(), dbProduct)
	if err != nil {
		log.Fatal("Error saving product", err)
		return &domain.Product{}, err
	}

	return mapProductToDomainModel(dbProduct), err
}

func mapProductToMongoModel(product *domain.Product) *ProductModel {
	return &ProductModel{
		ID: product.ID, Name: product.Name, Description: product.Description, Category: product.Category,
		Price: product.Price, Stock: product.Stock, IDSeller: product.IDSeller, CreatedAt: time.Now(),
	}
}

func mapProductToDomainModel(product *ProductModel) *domain.Product {
	return domain.NewProduct(product.ID, product.Name, product.Description, product.Category, product.Price,
		product.Stock, product.IDSeller)
}
