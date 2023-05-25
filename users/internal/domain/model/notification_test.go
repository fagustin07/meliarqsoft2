package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidNotification(t *testing.T) {
	notification, err := NewNotification("Mauro", "mauro@fede.com")

	assert.Equal(t, notification.Email, "mauro@fede.com")
	assert.NoError(t, err)
}

func Test_CannotCreateANotificationWithInvalidAddress(t *testing.T) {
	notification, err := NewNotification("Mauro", "mail invalido")

	assert.Nil(t, notification)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid email address")
}
