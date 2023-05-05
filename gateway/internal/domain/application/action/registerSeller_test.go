package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_RegisterSellerEvent(t *testing.T) {
	registerSellerEvent, mocks := setUpRegisterSeller(t)

	id, _ := uuid.NewUUID()
	seller := &model.Seller{}

	mocks.SellerRepository.EXPECT().Create(seller).Return(id, nil)

	newId, err := registerSellerEvent.Execute(seller)

	assert.Equal(t, id, newId)
	assert.NoError(t, err)
}

func Test_RegisterSellerEvent_ThrowsError(t *testing.T) {
	registerSellerEvent, mocks := setUpRegisterSeller(t)

	seller := &model.Seller{}

	mocks.SellerRepository.EXPECT().Create(seller).Return(uuid.Nil, errors.New(""))

	blankID, err := registerSellerEvent.Execute(seller)

	assert.Equal(t, blankID, uuid.Nil)
	assert.Error(t, err)
}

func setUpRegisterSeller(t *testing.T) (*RegisterSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewRegisterSellerEvent(mocks.SellerRepository), mocks
}
