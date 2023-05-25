package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidStock(t *testing.T) {
	newStock, err := NewStock(5)

	assert.NoError(t, err)
	assert.Equal(t, newStock.Amount, 5)
}

func Test_CannotCreateNegativeStock(t *testing.T) {
	newStock, err := NewStock(-5)

	assert.Error(t, err)
	assert.Nil(t, newStock)
	assert.Equal(t, err.Error(), "invalid stock amount")
}
