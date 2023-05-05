package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {

		return err
	}

	return nil
}
