package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindProductsBySeller(t *testing.T) {
	findProductsBySeller, mocks := setUpFindProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()

	prod := model.Product{ID: idProd, IDSeller: idSeller}
	products := []model.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(products, nil)

	prodsRes, err := findProductsBySeller.Execute(idSeller)

	assert.NoError(t, err)
	assert.Equal(t, prodsRes[0].IDSeller, idSeller)
	assert.Len(t, prodsRes, 1)
}

func setUpFindProductsBySeller(t *testing.T) (*FindProductsBySeller, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindProductsBySellerCommand(mocks.ProductRepository), mocks
}
