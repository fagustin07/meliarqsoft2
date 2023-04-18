package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindByIdProductWithResults(t *testing.T) {
	filterProdEvent, mocks := setUpFindByIdProdEvent(t)

	id, _ := uuid.NewUUID()

	prod := model.Product{ID: id}
	mocks.ProductRepository.EXPECT().FindById(id).Return(&prod, nil)

	resProd, err := filterProdEvent.Execute(id)

	assert.NoError(t, err)
	assert.Equal(t, prod.ID, resProd.ID)
}

func setUpFindByIdProdEvent(t *testing.T) (*FindProductById, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	return NewFindProductCommand(mocks.ProductRepository), mocks
}
