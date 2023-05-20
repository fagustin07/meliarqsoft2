package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_ExistCustomer(t *testing.T) {
	existCustomer, mocks := setUpExistCustomer(t)

	id, _ := uuid.NewUUID()
	mocks.CustomerRepository.EXPECT().FindById(id).Return(&model.Customer{}, nil)

	err := existCustomer.Execute(id)

	assert.NoError(t, err)
}

func Test_ExistCustomerThrowsErrorWhenNotFoundCustomer(t *testing.T) {
	existCustomer, mocks := setUpExistCustomer(t)

	id, _ := uuid.NewUUID()
	mocks.CustomerRepository.EXPECT().FindById(id).Return(&model.Customer{}, errors.New(""))

	err := existCustomer.Execute(id)

	assert.Error(t, err)
}

func setUpExistCustomer(t *testing.T) (*ExistCustomer, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewExistCustomerCommand(mocks.CustomerRepository), mocks
}
