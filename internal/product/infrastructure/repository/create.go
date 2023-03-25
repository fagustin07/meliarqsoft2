package repository

import (
	"context"
	"github.com/google/uuid"
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

func (repo ProductMongoDBRepository) Update(ID uuid.UUID, name string, description string, price float32, stock int, idSeller int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductMongoDBRepository) Delete(ID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (repo ProductMongoDBRepository) Find(name string, category string) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductMongoDBRepository) Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
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
