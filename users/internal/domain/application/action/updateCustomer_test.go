package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UpdateCustomerEvent(t *testing.T) {
	updateCustomer, mocks := setUpUpdateCustomer(t)

	id, _ := uuid.NewUUID()
	mocks.CustomerRepository.EXPECT().Update(id, "chester", "sandoval", "chester@gmail.com")

	assert.NoError(t, updateCustomer.Execute(id, "chester", "sandoval", "chester@gmail.com"))
}

func Test_UpdateCustomerEventOnInvalidEmailGivenThrowsError(t *testing.T) {
	updateUser, _ := setUpUpdateCustomer(t)
	id, _ := uuid.NewUUID()

	assert.Error(t, updateUser.Execute(id, "chester", "sandoval", "chester"))
}

func setUpUpdateCustomer(t *testing.T) (*UpdateCustomerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUpdateCustomerEvent(mocks.CustomerRepository), mocks
}
