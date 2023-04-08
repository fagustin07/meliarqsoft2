package dto

type CreateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required,email"`
}

type UpdateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required,email"`
}
