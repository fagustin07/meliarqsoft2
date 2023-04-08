package dto

type CreateUserRequest struct {
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Email   string `json:"email" bson:"email"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Email   string `json:"email" bson:"email"`
}
