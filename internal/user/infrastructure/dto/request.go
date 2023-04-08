package dto

type CreateUserRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}
