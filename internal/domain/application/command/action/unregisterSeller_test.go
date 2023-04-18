package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain"
	query2 "meliarqsoft2/internal/domain/application/command/query"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UnregisterSellerWithProducts(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	prod := domain.Product{ID: idProd}
	prods := []domain.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(prods, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(nil, nil)

	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&domain.Seller{}, nil)
	mocks.SellerRepository.EXPECT().Delete(idSeller).Return(nil)

	err := unregisterSeller.Execute(idSeller)

	assert.NoError(t, err)
}

func setUpUnregisterSeller(t *testing.T) (*UnregisterSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	findPurchase := query2.NewFindPurchasesFromProductEvent(mocks.PurchaseRepository)
	undoPurchaseCommand := NewUndoPurchaseCommand(mocks.PurchaseRepository)
	undoPurchaseByProd := NewUndoPurchasesByProductCommand(findPurchase, undoPurchaseCommand)

	deleteProduct := NewDeleteProductEvent(mocks.ProductRepository, undoPurchaseByProd)
	findProductsBySeller := query2.NewFindProductsBySellerCommand(mocks.ProductRepository)

	deleteProdBySeller := NewDeleteProductsBySellerCommand(mocks.ProductRepository, findProductsBySeller, deleteProduct)

	return NewUnregisterSellerEvent(
		mocks.SellerRepository,
		deleteProdBySeller,
		query2.NewExistSellerCommand(mocks.SellerRepository),
	), mocks
}
