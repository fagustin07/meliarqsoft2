package model

type SellerNotFoundError struct {
	Id string
}

func (s SellerNotFoundError) Error() string {
	return "seller " + s.Id + " not found"
}

type SellerAlreadyExist struct{}

func (s SellerAlreadyExist) Error() string {
	return "email or business name already taken"
}
