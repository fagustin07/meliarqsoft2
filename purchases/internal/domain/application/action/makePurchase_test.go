package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_MakePurchaseSuccessfully(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	ProductName := "asdas"
	SellerName := "asdasd"
	SellerEmail := "prueba@gmail.com"
	BuyerName := "asdasd"
	BuyerEmail := "prueba@gmail.com"
	idPurch, _ := uuid.NewUUID()
	units := 10

	purchaseReq, err := dto.CreatePurchaseRequest{
		IDProduct:   idProd,
		IDUser:      idUser,
		ProductName: ProductName,
		SellerName:  SellerName,
		SellerEmail: SellerEmail,
		BuyerName:   BuyerName,
		BuyerEmail:  BuyerEmail,
		Units:       units,
		Total:       32.2,
	}.MapToInternal()

	log.Println(err)
	mocks.PurchaseRepository.EXPECT().Create(purchaseReq.Purchase).Return(idPurch, nil)
	mocks.NotificationRepository.EXPECT().Send(purchaseReq.SellerNotification).Return(nil)
	mocks.NotificationRepository.EXPECT().Send(purchaseReq.BuyerNotification).Return(nil)

	newPurchase, err := makePurchaseEvent.Execute(purchaseReq)

	assert.NoError(t, err)
	assert.Equal(t, newPurchase, purchaseReq.Purchase)
}

func Test_MakePurchaseReturnErrorWhenFailOnCreatePurchase(t *testing.T) {
	makePurchaseEvent, mocks := setUpMakePurchase(t)
	idProd, _ := uuid.NewUUID()
	idUser, _ := uuid.NewUUID()
	ProductName := "asdas"
	SellerName := "asdasd"
	SellerEmail := "prueba@gmail.com"
	BuyerName := "asdasd"
	BuyerEmail := "prueba@gmail.com"
	units := 10

	purchase, _ := dto.CreatePurchaseRequest{
		IDProduct:   idProd,
		IDUser:      idUser,
		ProductName: ProductName,
		SellerName:  SellerName,
		SellerEmail: SellerEmail,
		BuyerName:   BuyerName,
		BuyerEmail:  BuyerEmail,
		Units:       units,
		Total:       32.2,
	}.MapToInternal()

	mocks.PurchaseRepository.EXPECT().Create(purchase.Purchase).Return(uuid.Nil, errors.New(""))
	mocks.NotificationRepository.EXPECT().Send(purchase.SellerNotification).Times(0)
	mocks.NotificationRepository.EXPECT().Send(purchase.BuyerNotification).Times(0)
	_, err := makePurchaseEvent.Execute(purchase)

	assert.Error(t, err)
}

func setUpMakePurchase(t *testing.T) (*MakePurchaseEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewMakePurchaseEvent(mocks.PurchaseRepository, mocks.NotificationRepository), mocks
}
