package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindByIdCustomer(t *testing.T) {
	findByIdCustomer, mocks := setUpFindByIdCustomer(t)
	id := uuid.New()
	email, _ := model.NewEmail("chester@gmail.com")

	customer := &model.Customer{
		ID:      id,
		Name:    "chester",
		Surname: "sandoval",
		Email:   email,
	}

	mocks.CustomerRepository.EXPECT().FindById(id).Return(customer, nil)

	customerFounded, err := findByIdCustomer.Execute(id)

	assert.Equal(t, customer, customerFounded)
	assert.NoError(t, err)
}

func Test_FindByIdCustomerThrowsErrorWhenNotFoundCustomer(t *testing.T) {
	findByIdCustomer, mocks := setUpFindByIdCustomer(t)

	id, _ := uuid.NewUUID()
	mocks.CustomerRepository.EXPECT().FindById(id).Return(&model.Customer{}, errors.New(""))

	_, err := findByIdCustomer.Execute(id)

	assert.Error(t, err)
}

func setUpFindByIdCustomer(t *testing.T) (*FindCustomerByIdEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindCustomerByIdEvent(mocks.CustomerRepository), mocks
}
