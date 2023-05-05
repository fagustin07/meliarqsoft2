package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_CreateProductEvent(t *testing.T) {
	createProdEvent, mocks := setUpCreateProdEvent(t)

	newUUID, _ := uuid.NewUUID()
	newProduct, _ := dto.CreateProductRequest{
		Name:        "DDL",
		Description: "dulce de lece",
		Category:    "dulces",
		Price:       150,
		Stock:       200,
		IDSeller:    uuid.Nil,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().Create(newProduct).Return(newUUID, nil)

	resultId, err := createProdEvent.Execute(newProduct)

	assert.Equal(t, resultId, newUUID)
	assert.NoError(t, err)
}

func setUpCreateProdEvent(t *testing.T) (*CreateProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewCreateProductEvent(mocks.ProductRepository), mocks
}
