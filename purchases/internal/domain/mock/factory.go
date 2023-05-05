package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type RepositoriesMock struct {
	PurchaseRepository *MockIPurchaseRepository
}

func NewMockRepositories(t *testing.T) *RepositoriesMock {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return &RepositoriesMock{
		PurchaseRepository: NewMockIPurchaseRepository(ctrl),
	}
}
