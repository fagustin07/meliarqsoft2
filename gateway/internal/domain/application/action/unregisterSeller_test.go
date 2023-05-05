package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_UnregisterSellerWithProducts(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	prod := model.Product{ID: idProd}
	prods := []model.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(prods, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(nil, nil)

	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, nil)
	mocks.SellerRepository.EXPECT().Delete(idSeller).Return(nil)

	err := unregisterSeller.Execute(idSeller)

	assert.NoError(t, err)
}

func Test_UnregisterSellerWithProductsThrowsErrorWhenDoesNotExistSeller(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idSeller, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, errors.New("")).Times(1)

	err := unregisterSeller.Execute(idSeller)

	assert.Error(t, err)
}

func Test_UnregisterSellerWithProductsThrowsErrorWhenDeletingProductsFails(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	prod := model.Product{ID: idProd}
	prods := []model.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(prods, errors.New(""))
	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, nil)

	err := unregisterSeller.Execute(idSeller)

	assert.Error(t, err)
}

func Test_UnregisterSellerWithProducts_ThrowsErrorWhenDeleteSellerFails(t *testing.T) {
	unregisterSeller, mocks := setUpUnregisterSeller(t)
	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	prod := model.Product{ID: idProd}
	prods := []model.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(prods, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(nil, nil)

	mocks.SellerRepository.EXPECT().FindById(idSeller).Return(&model.Seller{}, nil)
	mocks.SellerRepository.EXPECT().Delete(idSeller).Return(errors.New(""))

	err := unregisterSeller.Execute(idSeller)

	assert.Error(t, err)
}

func setUpUnregisterSeller(t *testing.T) (*UnregisterSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	findPurchase := query.NewFindPurchasesFromProductEvent(mocks.PurchaseRepository)
	undoPurchaseCommand := NewUndoPurchaseCommand(mocks.PurchaseRepository)
	undoPurchaseByProd := NewUndoPurchasesByProductCommand(findPurchase, undoPurchaseCommand)

	deleteProduct := NewDeleteProductEvent(mocks.ProductRepository, undoPurchaseByProd)
	findProductsBySeller := query.NewFindProductsBySellerCommand(mocks.ProductRepository)

	deleteProdBySeller := NewDeleteProductsBySellerCommand(mocks.ProductRepository, findProductsBySeller, deleteProduct)

	return NewUnregisterSellerEvent(
		mocks.SellerRepository,
		deleteProdBySeller,
		query.NewExistSellerCommand(mocks.SellerRepository),
	), mocks
}
