package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type RepositoriesMock struct {
	CustomerRepository *MockICustomerRepository
	SellerRepository   *MockISellerRepository
}

func NewMockRepositories(t *testing.T) *RepositoriesMock {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return &RepositoriesMock{
		CustomerRepository: NewMockICustomerRepository(ctrl),
		SellerRepository:   NewMockISellerRepository(ctrl),
	}
}
