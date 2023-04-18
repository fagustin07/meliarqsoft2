package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type RepositoriesMock struct {
	UserRepository     *MockIUserRepository
	SellerRepository   *MockISellerRepository
	ProductRepository  *MockIProductRepository
	PurchaseRepository *MockIPurchaseRepository
}

func NewMockRepositories(t *testing.T) *RepositoriesMock {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return &RepositoriesMock{
		UserRepository:     NewMockIUserRepository(ctrl),
		SellerRepository:   NewMockISellerRepository(ctrl),
		ProductRepository:  NewMockIProductRepository(ctrl),
		PurchaseRepository: NewMockIPurchaseRepository(ctrl),
	}
}
