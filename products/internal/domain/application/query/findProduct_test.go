package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindProductWithResults(t *testing.T) {
	filterProdEvent, mocks := setUpFindProdEvent(t)

	id, _ := uuid.NewUUID()

	prod := model.Product{ID: id}
	products := []model.Product{prod}

	name := "fa"
	category := "cultad"
	mocks.ProductRepository.EXPECT().Find(name, category).Return(products, nil)

	resProds, err := filterProdEvent.Execute(name, category)

	assert.NoError(t, err)
	assert.Equal(t, resProds[0].ID, prod.ID)
}

func Test_FindProduct_ThrowsErrorWhenMinPriceIsGreatherThanMaxPrice(t *testing.T) {
	filterProdEvent, mocks := setUpFindProdEvent(t)

	name := "fa"
	category := "cultad"
	mocks.ProductRepository.EXPECT().Find(name, category).Return(nil, nil)

	resProds, _ := filterProdEvent.Execute(name, category)

	assert.Nil(t, resProds)
}

func Test_FindProduct_ThrowsErrorWhenFindFails(t *testing.T) {
	filterProdEvent, mocks := setUpFindProdEvent(t)

	name := "fa"
	category := "cultad"
	mocks.ProductRepository.EXPECT().Find(name, category).Return(nil, errors.New(""))

	_, err := filterProdEvent.Execute(name, category)

	assert.Error(t, err)
}

func setUpFindProdEvent(t *testing.T) (*FindProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	return NewFindProductEvent(mocks.ProductRepository), mocks
}
