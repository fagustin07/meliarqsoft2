package dto

type CreateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name"`
	Email        string `json:"email" bson:"email"`
}

type UpdateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name"`
	Email        string `json:"email" bson:"email"`
}
