package seller

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	res, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("cannot delete user " + ID.String())
	}

	return nil
}
