package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidEmail(t *testing.T) {
	email, err := NewEmail("fede@sandoval.com")

	assert.Equal(t, email.Address, "fede@sandoval.com")
	assert.NoError(t, err)
}

func Test_CannotCreateAnEmailWithInvalidAddress(t *testing.T) {
	email, err := NewEmail("federico")

	assert.Nil(t, email)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid email address")
}
