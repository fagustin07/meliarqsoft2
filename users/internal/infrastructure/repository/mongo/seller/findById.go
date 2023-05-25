package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) FindById(idSeller uuid.UUID) (*model.Seller, error) {
	filter := bson.M{
		"_id":        idSeller,
		"deleted_at": time.Time{},
	}

	var sellerDb SellerModel
	err := repo.collection.FindOne(context.Background(), filter).Decode(&sellerDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.SellerNotFoundError{Id: idSeller.String()}
		}
		return nil, err
	}

	return MapSellerFromModel(&sellerDb)
}
