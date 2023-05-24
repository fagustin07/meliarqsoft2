package purchase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository MongoRepository) DeleteAll() error {
	_, err := repository.collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	return nil
}
