package dto

type FilterProductRequest struct {
	MinPrice float32 `json:"min_price" bson:"min_price"`
	MaxPrice float32 `json:"max_price" bson:"max_price"`
}
