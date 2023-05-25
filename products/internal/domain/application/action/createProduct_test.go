package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_CreateProductEvent(t *testing.T) {
	createProdEvent, mocks := setUpCreateProdEvent(t)

	newUUID, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	newProduct, _ := dto.CreateProductRequest{
		Name:        "DDL",
		Description: "dulce de lece",
		Category:    "dulces",
		Price:       150,
		Stock:       200,
		IDSeller:    idSeller,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().Create(newProduct).Return(newUUID, nil)
	mocks.FindSellerId.EXPECT().Execute(idSeller).Return(model.SellerDTO{}, nil)

	resultId, err := createProdEvent.Execute(newProduct)

	assert.Equal(t, resultId, newUUID)
	assert.NoError(t, err)
}

func Test_CreateProductEventFailsWhenSellerDoesNotExist(t *testing.T) {
	createProdEvent, mocks := setUpCreateProdEvent(t)

	newUUID, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	newProduct, _ := dto.CreateProductRequest{
		Name:        "DDL",
		Description: "dulce de lece",
		Category:    "dulces",
		Price:       150,
		Stock:       200,
		IDSeller:    idSeller,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().Create(newProduct).Return(newUUID, nil)
	mocks.FindSellerId.EXPECT().Execute(idSeller).Return(model.SellerDTO{}, errors.New(""))

	_, err := createProdEvent.Execute(newProduct)

	assert.Error(t, err)
}

func setUpCreateProdEvent(t *testing.T) (*CreateProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewCreateProductEvent(mocks.ProductRepository, mocks.FindSellerId), mocks
}
