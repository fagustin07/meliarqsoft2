package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/application/command/query"
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

	mocks.SellerRepository.EXPECT().FindById(uuid.Nil).Return(nil, nil)
	mocks.ProductRepository.EXPECT().Create(newProduct).Return(newUUID, nil)

	resultId, err := createProdEvent.Execute(newProduct)

	assert.Equal(t, resultId, newUUID)
	assert.NoError(t, err)
}

func Test_CreateProductEvent_FailsWhenSellerDoesNotExist(t *testing.T) {
	createProdEvent, mocks := setUpCreateProdEvent(t)

	newProduct, _ := dto.CreateProductRequest{
		Name:        "DDL",
		Description: "dulce de leche",
		Category:    "dulces",
		Price:       150,
		Stock:       200,
		IDSeller:    uuid.Nil,
	}.MapToModel()

	mocks.SellerRepository.EXPECT().FindById(uuid.Nil).Return(nil, errors.New(""))

	resultId, err := createProdEvent.Execute(newProduct)

	assert.Equal(t, resultId, uuid.Nil)
	assert.Error(t, err)
}

func setUpCreateProdEvent(t *testing.T) (*CreateProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	existSeller := query.NewExistSellerCommand(mocks.SellerRepository)

	return NewCreateProductEvent(mocks.ProductRepository, existSeller), mocks
}
