package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type RepositoriesMock struct {
	ProductRepository *MockIProductRepository
	FindSellerId      *MockIFindSellerByIdService
	DeletePurchases   *MockIDeletePurchasesByProductsService
}

func NewMockRepositories(t *testing.T) *RepositoriesMock {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return &RepositoriesMock{
		ProductRepository: NewMockIProductRepository(ctrl),
		FindSellerId:      NewMockIFindSellerByIdService(ctrl),
		DeletePurchases:   NewMockIDeletePurchasesByProductsService(ctrl),
	}
}
