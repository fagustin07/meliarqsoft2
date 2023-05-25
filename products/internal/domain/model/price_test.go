package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidPrice(t *testing.T) {
	newPrice, err := NewPrice(float32(5))

	assert.NoError(t, err)
	assert.Equal(t, newPrice.Value, float32(5))
}

func Test_CannotCreateNegativePrice(t *testing.T) {
	newPrice, err := NewPrice(float32(-5))

	assert.Error(t, err)
	assert.Nil(t, newPrice)
	assert.Equal(t, err.Error(), "invalid price value")
}
