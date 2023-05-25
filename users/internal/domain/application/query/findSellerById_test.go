package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindByIdSeller(t *testing.T) {
	findByIdSeller, mocks := setUpFindByIdSeller(t)
	id := uuid.New()
	email, _ := model.NewEmail("chester@gmail.com")

	customer := &model.Seller{
		ID:           id,
		BusinessName: "chester SRL",
		Email:        email,
	}

	mocks.SellerRepository.EXPECT().FindById(id).Return(customer, nil)

	customerFounded, err := findByIdSeller.Execute(id)

	assert.Equal(t, customer, customerFounded)
	assert.NoError(t, err)
}

func Test_FindByIdSellerThrowsErrorWhenNotFoundSeller(t *testing.T) {
	findByIdSeller, mocks := setUpFindByIdSeller(t)

	id, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().FindById(id).Return(&model.Seller{}, errors.New(""))

	_, err := findByIdSeller.Execute(id)

	assert.Error(t, err)
}

func setUpFindByIdSeller(t *testing.T) (*FindSellerByIdEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindSellerByIdEvent(mocks.SellerRepository), mocks
}
