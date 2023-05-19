package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_FindCustomerEvent(t *testing.T) {
	findCustomerEvent, mocks := setUpFindCustomerEvent(t)

	id, err := uuid.NewUUID()

	customer, _ := dto.CustomerDTO{
		ID:      id,
		Surname: "sanchez",
		Name:    "chester",
		Email:   "chester99@gmail.com",
	}.MapToModel()
	customersList := []*model.Customer{&customer}

	mocks.CustomerRepository.EXPECT().Find("99").Return(customersList, nil)

	customersRes, err := findCustomerEvent.Execute("99")

	assert.NoError(t, err)
	assert.Len(t, customersRes, 1)
	assert.Equal(t, customersRes[0].ID, customer.ID)
	assert.Equal(t, customersRes[0].Name, customer.Name)
	assert.Equal(t, customersRes[0].Surname, customer.Surname)
	assert.Equal(t, customersRes[0].Email, customer.Email)

}

func setUpFindCustomerEvent(t *testing.T) (*FindCustomerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindCustomerEvent(mocks.CustomerRepository), mocks
}
