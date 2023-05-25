package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FilterProductWithResults(t *testing.T) {
	filterProdEvent, mocks := setUpFilterProdEvent(t)

	id, _ := uuid.NewUUID()

	prod := model.Product{ID: id}
	products := []model.Product{prod}

	minPrice := float32(40)
	maxPrice := float32(40)
	mocks.ProductRepository.EXPECT().Filter(minPrice, maxPrice).Return(products, nil)

	resProds, err := filterProdEvent.Execute(minPrice, maxPrice)

	assert.NoError(t, err)
	assert.Equal(t, resProds[0].ID, prod.ID)
}

func Test_FilterProduct_ThrowsErrorWhenMinPriceIsGreatherThanMaxPrice(t *testing.T) {
	filterProdEvent, _ := setUpFilterProdEvent(t)

	maxPrice := float32(40)
	minPrice := maxPrice + float32(80)

	_, err := filterProdEvent.Execute(minPrice, maxPrice)

	assert.Error(t, err)
}

func setUpFilterProdEvent(t *testing.T) (*FilterProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	return NewFilterProductEvent(mocks.ProductRepository), mocks
}
