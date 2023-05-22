package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_UnregisterSeller(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idSeller, _ := uuid.NewUUID()

	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, nil)
	mocks.SellerRepository.EXPECT().Delete(idSeller).Return(nil)

	err := unregisterSeller.Execute(idSeller)

	assert.NoError(t, err)
}

func Test_UnregisterSeller_ThrowsErrorWhenDeleteSellerFails(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idSeller, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, nil)
	mocks.SellerRepository.EXPECT().Delete(idSeller).Return(errors.New(""))

	err := unregisterSeller.Execute(idSeller)

	assert.Error(t, err)
}

func setUpUnregisterSeller(t *testing.T) (*UnregisterSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUnregisterSellerEvent(mocks.SellerRepository), mocks
}
