package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_MakePurchaseSuccessfully(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	ProductName := ""
	SellerName := ""
	SellerEmail := ""
	BuyerName := ""
	BuyerEmail := ""
	idPurch, _ := uuid.NewUUID()
	units := 10

	purchase, _, _, _ := dto.CreatePurchaseRequest{
		IDProduct:   idProd,
		IDUser:      idUser,
		ProductName: ProductName,
		SellerName:  SellerName,
		SellerEmail: SellerEmail,
		BuyerName:   BuyerName,
		BuyerEmail:  BuyerEmail,
		Units:       units,
		Total:       32.2,
	}.MapToModel()

	mocks.PurchaseRepository.EXPECT().Create(&purchase).Return(idPurch, nil)

	newId, err := makePurchaseEvent.Execute(&purchase)

	assert.NoError(t, err)
	assert.Equal(t, newId, idPurch)
}

func Test_MakePurchaseReturnErrorWhenFailOnCreatePurchase(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	ProductName := ""
	SellerName := ""
	SellerEmail := ""
	BuyerName := ""
	BuyerEmail := ""
	units := 10

	purchase, _, _, _ := dto.CreatePurchaseRequest{
		IDProduct:   idProd,
		IDUser:      idUser,
		ProductName: ProductName,
		SellerName:  SellerName,
		SellerEmail: SellerEmail,
		BuyerName:   BuyerName,
		BuyerEmail:  BuyerEmail,
		Units:       units,
		Total:       32.2,
	}.MapToModel()

	mocks.PurchaseRepository.EXPECT().Create(&purchase).Return(uuid.Nil, errors.New(""))

	newId, err := makePurchaseEvent.Execute(&purchase)

	assert.Error(t, err)
	assert.Equal(t, newId, uuid.Nil)
}

func setUpMakePurchase(t *testing.T) (*MakePurchaseEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewMakePurchaseEvent(mocks.PurchaseRepository), mocks
}
