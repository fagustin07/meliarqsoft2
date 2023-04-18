package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain"
	query2 "meliarqsoft2/internal/domain/application/command/query"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_DeleteProductsBySellerWithoutProducts(t *testing.T) {
	deleteBySellerCommand, mocks := setUpDeleteProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(nil, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&domain.Product{}, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(nil, nil)

	err := deleteBySellerCommand.Execute(idSeller)

	assert.NoError(t, err)
}

func Test_DeleteProductsBySellerWithProducts(t *testing.T) {
	deleteBySellerCommand, mocks := setUpDeleteProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	prod := domain.Product{ID: idProd}
	prods := []domain.Product{prod}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(prods, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.PurchaseRepository.EXPECT().Find(idProd).Return(nil, nil)

	err := deleteBySellerCommand.Execute(idSeller)

	assert.NoError(t, err)
}

func setUpDeleteProductsBySeller(t *testing.T) (*DeleteProductsBySeller, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	findPurchase := query2.NewFindPurchasesFromProductEvent(mocks.PurchaseRepository)
	undoPurchaseCommand := NewUndoPurchaseCommand(mocks.PurchaseRepository)
	undoPurchaseByProd := NewUndoPurchasesByProductCommand(findPurchase, undoPurchaseCommand)

	deleteProduct := NewDeleteProductEvent(mocks.ProductRepository, undoPurchaseByProd)
	findProductsBySeller := query2.NewFindProductsBySellerCommand(mocks.ProductRepository)

	return NewDeleteProductsBySellerCommand(mocks.ProductRepository, findProductsBySeller, deleteProduct), mocks
}
