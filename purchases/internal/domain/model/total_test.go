package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidTotal(t *testing.T) {
	newTotal, err := NewTotal(float32(5))

	assert.NoError(t, err)
	assert.Equal(t, newTotal.Value, float32(5))
}

func Test_CannotCreateNegativeTotal(t *testing.T) {
	newTotal, err := NewTotal(float32(-5))

	assert.Error(t, err)
	assert.Nil(t, newTotal)
	assert.Equal(t, err.Error(), "invalid total price amount")
}
