package domain

type Purchase struct {
	IDProduct int
	IDUser    int
}

func NewPurchase(IDProduct int, IDUser int) *Purchase {
	return &Purchase{IDProduct: IDProduct, IDUser: IDUser}
}
