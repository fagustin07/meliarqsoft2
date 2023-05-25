package seller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo MongoRepository) DeleteAll() error {
	_, err := repo.collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	return nil
}
