package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (model.Product, error) {
	_, err := repo.FindById(ID)

	if err != nil {
		return model.Product{}, model2.ProductNotFound{Id: ID.String()}
	}

	var fieldsToUpdate bson.D
	if name != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "name", Value: name})
	}

	if description != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "description", Value: description})
	}

	if category != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "category", Value: category})
	}

	if price > 0 {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "price", Value: price})
	}

	if stock > 0 {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "stock", Value: stock})
	}

	_, err = repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID, "deleted_at": time.Time{}},
		bson.D{{"$set", fieldsToUpdate}},
	)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Product{}, model2.ProductNotFound{Id: ID.String()}
		}
		if mongo.IsDuplicateKeyError(err) {
			return model.Product{}, model2.ProductAlreadyExist{}
		}
		return model.Product{}, err
	}

	return model.Product{}, nil
}
