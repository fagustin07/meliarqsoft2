package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (model.Product, error) {
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

	x, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", fieldsToUpdate}},
	)
	log.Println(x)

	if err != nil {
		log.Print(err)
		return model.Product{}, err
	}

	return model.Product{}, nil
}
