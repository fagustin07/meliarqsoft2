package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo MongoDBRepository) Delete(ID uuid.UUID) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
