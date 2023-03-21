package application

type ISellerService interface {
	Create(businessName string, email string) error
}
