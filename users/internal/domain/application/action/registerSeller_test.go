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
	noti, _ := model.NewNotification("", "fede@gmail.com")

	mocks.SellerRepository.EXPECT().Create(seller).Return(id, nil)
	mocks.NotificationRepository.EXPECT().Send(noti).Return(nil)

	newId, err := registerSellerEvent.Execute(seller, noti)

	assert.Equal(t, id, newId)
	assert.NoError(t, err)
}

func Test_RegisterSellerEvent_ThrowsError(t *testing.T) {
	registerSellerEvent, mocks := setUpRegisterSeller(t)

	seller := &model.Seller{}
	noti, _ := model.NewNotification("", "fede@gmail.com")

	mocks.SellerRepository.EXPECT().Create(seller).Return(uuid.Nil, errors.New(""))
	mocks.NotificationRepository.EXPECT().Send(noti).Times(0)

	blankID, err := registerSellerEvent.Execute(seller, noti)

	assert.Equal(t, blankID, uuid.Nil)
	assert.Error(t, err)
}

func setUpRegisterSeller(t *testing.T) (*RegisterSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewRegisterSellerEvent(mocks.SellerRepository, mocks.NotificationRepository), mocks
}
