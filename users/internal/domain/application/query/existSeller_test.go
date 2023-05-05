package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_ExistSeller(t *testing.T) {
	existSeller, mocks := setUpExistSeller(t)

	id, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().FindById(id).Return(&model.Seller{}, nil)

	err := existSeller.Execute(id)

	assert.NoError(t, err)
}

func Test_ExistSellerThrowsErrorWhenNotFoundSeller(t *testing.T) {
	existSeller, mocks := setUpExistSeller(t)

	id, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().FindById(id).Return(&model.Seller{}, errors.New(""))

	err := existSeller.Execute(id)

	assert.Error(t, err)
}

func setUpExistSeller(t *testing.T) (*ExistSeller, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewExistSellerCommand(mocks.SellerRepository), mocks
}
