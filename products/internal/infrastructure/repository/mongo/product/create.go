package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) Create(product model.Product) (uuid.UUID, error) {
	var newUUID uuid.UUID
	var err error

	if product.ID == uuid.Nil {
		newUUID, err = uuid.NewUUID()
		if err != nil {
			return uuid.Nil, model2.CreateUUIDError{}
		}
	} else {
		newUUID = product.ID
	}

	product.ID = newUUID

	dbProduct := mapProductToMongoModel(product)
	_, err = repo.collection.InsertOne(context.Background(), dbProduct)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return uuid.Nil, model2.ProductAlreadyExist{}
		}
		return uuid.Nil, err
	}

	return newUUID, nil
}

func mapProductToMongoModel(product model.Product) *ProductModel {
	return &ProductModel{
		ID: product.ID, Name: product.Name, Description: product.Description, Category: product.Category,
		Price: product.Price.Value, Stock: product.Stock.Amount, IDSeller: product.IDSeller,
		CreatedAt: time.Now(),
	}
}

func mapProductToDomainModel(product *ProductModel) (model.Product, error) {
	newPrice, err := model.NewPrice(product.Price)
	if err != nil {
		return model.Product{}, err
	}

	stock, err := model.NewStock(product.Stock)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       newPrice,
		Stock:       stock,
		IDSeller:    product.IDSeller,
	}, nil
}
