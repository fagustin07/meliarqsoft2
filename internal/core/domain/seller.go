package domain

type Seller struct {
	ID           int
	BusinessName string
	Email        string
}

func NewSeller(id int, businessName string, email string) *Seller {
	return &Seller{ID: id, BusinessName: businessName, Email: email}
}
