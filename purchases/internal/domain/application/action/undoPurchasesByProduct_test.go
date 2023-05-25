package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_UndoPurchasesByProduct(t *testing.T) {
	undoPurchasesByProduct, mocks := setUpUndoPurchasesByProduct(t)

	idProd, _ := uuid.NewUUID()
	idPurch, _ := uuid.NewUUID()

	purch := &model.Purchase{ID: idPurch}
	purchases := []*model.Purchase{purch}
	ids := []uuid.UUID{idProd}

	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(purchases, nil)
	mocks.PurchaseRepository.EXPECT().DeleteByProductsIDs(ids).Return(nil)

	err := undoPurchasesByProduct.Execute(ids)

	assert.NoError(t, err)
}

func Test_UndoPurchasesByProduct_ThrowsErrorWhenUndoPurchaseFails(t *testing.T) {
	undoPurchasesByProduct, mocks := setUpUndoPurchasesByProduct(t)

	idProd, _ := uuid.NewUUID()
	idPurch, _ := uuid.NewUUID()
	ids := []uuid.UUID{idProd}

	purch := &model.Purchase{ID: idPurch}
	purchases := []*model.Purchase{purch}

	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(purchases, nil)
	mocks.PurchaseRepository.EXPECT().DeleteByProductsIDs(ids).Return(errors.New(""))

	err := undoPurchasesByProduct.Execute(ids)

	assert.Error(t, err)
}

func setUpUndoPurchasesByProduct(t *testing.T) (*UndoPurchasesByProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUndoPurchasesByProductEvent(mocks.PurchaseRepository), mocks
}
