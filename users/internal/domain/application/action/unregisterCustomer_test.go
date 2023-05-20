package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UnregisterCustomer(t *testing.T) {
	unregisterCustomer, mocks := setUpUnregisterCustomer(t)

	id, _ := uuid.NewUUID()
	mocks.CustomerRepository.EXPECT().Delete(id).Return(nil)

	err := unregisterCustomer.Execute(id)

	assert.NoError(t, err)
}

func setUpUnregisterCustomer(t *testing.T) (*UnregisterCustomerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	return NewUnregisterCustomerEvent(mocks.CustomerRepository), mocks
}
