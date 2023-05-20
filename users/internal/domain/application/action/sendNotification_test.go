package action

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_SendNotificationEvent(t *testing.T) {
	notificationSellerEvent, mocks := setUpSendNotification(t)

	notification, _ := model.NewNotification("", "test@gmail.com")

	mocks.NotificationRepository.EXPECT().Send(notification).Return(nil)

	err := notificationSellerEvent.Execute(notification.Name, notification.Email)

	assert.NoError(t, err)
}

func Test_SendNotificationEvent_ThrowsError(t *testing.T) {
	notificationSellerEvent, mocks := setUpSendNotification(t)

	notification := &model.Notification{}

	mocks.NotificationRepository.EXPECT().Send(notification).Return(errors.New(""))

	err := notificationSellerEvent.Execute("", "test")

	assert.Error(t, err)
}

func setUpSendNotification(t *testing.T) (*SendNotificationEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewSendNotificationEvent(mocks.NotificationRepository), mocks
}
