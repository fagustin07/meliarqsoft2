package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_MakePurchaseSuccessfully(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	idPurch, _ := uuid.NewUUID()

	stockBeforePurchase := 50
	unitsToConsume := 10
	product, _ := dto.ProductDTO{
		ID:          idProd,
		Name:        "chicha",
		Description: "rito",
		Category:    "hern",
		Price:       100,
		Stock:       stockBeforePurchase,
		IDSeller:    idSeller,
	}.MapToModel()

	purchase, _ := dto.CreatePurchaseRequest{
		IDProduct: idProd,
		IDUser:    idUser,
		Units:     unitsToConsume,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&product, nil).MaxTimes(2)
	mocks.ProductRepository.EXPECT().UpdateStock(idProd, stockBeforePurchase-unitsToConsume).Return(nil)
	mocks.UserRepository.EXPECT().FindById(idUser).Return(&model.User{}, nil)
	mocks.PurchaseRepository.EXPECT().Create(&purchase, &product).Return(idPurch, product.Price.Value*float32(unitsToConsume), nil)

	newId, totalPrice, err := makePurchaseEvent.Execute(&purchase)

	assert.NoError(t, err)
	assert.Equal(t, newId, idPurch)
	assert.Equal(t, totalPrice, product.Price.Value*float32(unitsToConsume))
}

func Test_MakePurchase_FailsWhenProductNotFound(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	unitsToConsume := 10
	purchase, _ := dto.CreatePurchaseRequest{
		IDProduct: idProd,
		IDUser:    idUser,
		Units:     unitsToConsume,
	}.MapToModel()
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&model.Product{}, errors.New(""))

	_, _, err := makePurchaseEvent.Execute(&purchase)

	assert.Error(t, err)
}

func Test_MakePurchaseFailsWhenUserDoesNotExist(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()

	stockBeforePurchase := 50
	unitsToConsume := 10
	product, _ := dto.ProductDTO{
		ID:          idProd,
		Name:        "chicha",
		Description: "rito",
		Category:    "hern",
		Price:       100,
		Stock:       stockBeforePurchase,
		IDSeller:    idSeller,
	}.MapToModel()
	purchase, _ := dto.CreatePurchaseRequest{
		IDProduct: idProd,
		IDUser:    idUser,
		Units:     unitsToConsume,
	}.MapToModel()
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&product, nil).MaxTimes(2)

	mocks.UserRepository.EXPECT().FindById(idUser).Return(&model.User{}, errors.New(""))

	_, _, err := makePurchaseEvent.Execute(&purchase)

	assert.Error(t, err)
}

func Test_MakePurchaseFails_WhenManageProductServiceFail(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	idPurch, _ := uuid.NewUUID()

	stockBeforePurchase := 50
	unitsToConsume := 10
	product, _ := dto.ProductDTO{
		ID:          idProd,
		Name:        "chicha",
		Description: "rito",
		Category:    "hern",
		Price:       100,
		Stock:       stockBeforePurchase,
		IDSeller:    idSeller,
	}.MapToModel()

	purchase, _ := dto.CreatePurchaseRequest{
		IDProduct: idProd,
		IDUser:    idUser,
		Units:     unitsToConsume,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&product, nil).Times(3)
	mocks.UserRepository.EXPECT().FindById(idUser).Return(&model.User{}, nil)
	mocks.ProductRepository.EXPECT().UpdateStock(idProd, stockBeforePurchase-unitsToConsume).Return(nil)
	mocks.ProductRepository.EXPECT().UpdateStock(idProd, stockBeforePurchase).Return(nil)
	mocks.PurchaseRepository.EXPECT().Create(&purchase, &product).Return(uuid.Nil, float32(0), errors.New(""))

	id, total, err := makePurchaseEvent.Execute(&purchase)

	assert.Error(t, err)
	assert.Equal(t, id, uuid.Nil)
	assert.NotEqual(t, id, idPurch)
	assert.Equal(t, total, float32(0))
}

func setUpMakePurchase(t *testing.T) (*MakePurchaseEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewMakePurchaseEvent(
		mocks.PurchaseRepository,
		query.NewFindProductCommand(mocks.ProductRepository),
		query.NewExistUserCommand(mocks.UserRepository),
		NewManageProductStockCommand(mocks.ProductRepository),
	), mocks
}
