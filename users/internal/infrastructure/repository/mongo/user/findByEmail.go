package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) FindByEmail(email string) (*model.User, error) {
	filter := bson.M{"email": email}
	var user model.User
	if err := repo.collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.UserNotFoundError{Id: ""}
		}
		return nil, err
	}
	return &user, nil
}
