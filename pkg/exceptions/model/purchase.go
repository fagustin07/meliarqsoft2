package model

type PurchaseNotFoundError struct {
	Id string
}

func (p PurchaseNotFoundError) Error() string {
	return "cannot found purchase " + p.Id
}

type InvalidUnitsError struct{}

func (i InvalidUnitsError) Error() string {
	return "invalid units amount"
}

type InvalidTotalPriceError struct{}

func (i InvalidTotalPriceError) Error() string {
	return "invalid total price amount"
}
