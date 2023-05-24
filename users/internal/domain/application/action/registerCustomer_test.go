package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_RegisterCustomerEvent(t *testing.T) {
	registerUserEvent, mocks := setUpRegisterCustomer(t)

	id, _ := uuid.NewUUID()
	user := &model.Customer{}
	noti, _ := model.NewNotification("", "fede@gmail.com")
	mocks.CustomerRepository.EXPECT().Create(user).Return(id, nil)
	mocks.NotificationRepository.EXPECT().Send(noti).Return(nil)

	newId, err := registerUserEvent.Execute(user, noti)

	assert.Equal(t, id, newId)
	assert.NoError(t, err)
}

func Test_RegisterCustomerEvent_ThrowsError(t *testing.T) {
	registerUserEvent, mocks := setUpRegisterCustomer(t)

	user := &model.Customer{}
	noti, _ := model.NewNotification("", "fede@gmail.com")

	mocks.CustomerRepository.EXPECT().Create(user).Return(uuid.Nil, errors.New(""))
	mocks.NotificationRepository.EXPECT().Send(noti).Times(0)

	blankID, err := registerUserEvent.Execute(user, nil)

	assert.Equal(t, blankID, uuid.Nil)
	assert.Error(t, err)
}

func setUpRegisterCustomer(t *testing.T) (*RegisterCustomerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewRegisterCustomerEvent(mocks.CustomerRepository, mocks.NotificationRepository), mocks
}
