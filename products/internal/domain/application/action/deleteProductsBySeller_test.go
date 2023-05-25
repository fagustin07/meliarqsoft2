package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_DeleteProductsBySeller(t *testing.T) {
	deleteBySellerCommand, mocks := setUpDeleteProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	ids := []uuid.UUID{idProd}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(nil, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&model.Product{}, nil)
	mocks.ProductRepository.EXPECT().DeleteBySeller(idSeller).Return(ids, nil)
	mocks.DeletePurchases.EXPECT().Execute(ids).Return(nil)

	deletedIds, err := deleteBySellerCommand.Execute(idSeller)

	assert.NoError(t, err)
	assert.Equal(t, ids, deletedIds)
}

func Test_DeleteProductsBySellerFailsWhenDeletePurchasesFailsAndRestoreProducts(t *testing.T) {
	deleteBySellerCommand, mocks := setUpDeleteProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()
	ids := []uuid.UUID{idProd}

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(nil, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&model.Product{}, nil)
	mocks.ProductRepository.EXPECT().DeleteBySeller(idSeller).Return(ids, nil)
	mocks.DeletePurchases.EXPECT().Execute(ids).Return(errors.New(""))
	mocks.ProductRepository.EXPECT().Restore(ids).Return(int64(len(ids)), nil)

	_, err := deleteBySellerCommand.Execute(idSeller)

	assert.Error(t, err)

}
func setUpDeleteProductsBySeller(t *testing.T) (*DeleteProductsBySeller, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewDeleteProductsBySellerEvent(mocks.ProductRepository, mocks.DeletePurchases), mocks
}
