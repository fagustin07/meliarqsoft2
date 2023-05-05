package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_FindPurchasesByProduct(t *testing.T) {
	findPurchasesByProduct, mocks := setUpFindPurchasesByProduct(t)

	idProd, _ := uuid.NewUUID()
	idPur, _ := uuid.NewUUID()

	purch := model.Purchase{ID: idPur, IDProduct: idProd}
	purchases := []*model.Purchase{&purch}

	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(purchases, nil)

	purchsRes, err := findPurchasesByProduct.Execute(idProd)

	assert.NoError(t, err)
	assert.Equal(t, purchsRes[0].IDProduct, idProd)
	assert.Equal(t, purchsRes[0].ID, purch.ID)
	assert.Len(t, purchsRes, 1)
}

func setUpFindPurchasesByProduct(t *testing.T) (*FindPurchasesFromProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindPurchasesFromProductEvent(mocks.PurchaseRepository), mocks
}
