package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateValidUnits(t *testing.T) {
	newUnits, err := NewUnits(5)

	assert.NoError(t, err)
	assert.Equal(t, newUnits.Amount, 5)
}

func Test_CannotCreateNegativeUnits(t *testing.T) {
	newUnits, err := NewUnits(-5)

	assert.Error(t, err)
	assert.Nil(t, newUnits)
	assert.Equal(t, err.Error(), "invalid units amount")
}

func Test_CannotCreateZeroUnits(t *testing.T) {
	newUnits, err := NewUnits(0)

	assert.Error(t, err)
	assert.Nil(t, newUnits)
	assert.Equal(t, err.Error(), "invalid units amount")
}
