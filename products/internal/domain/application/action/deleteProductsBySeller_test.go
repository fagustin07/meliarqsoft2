package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_DeleteProductsBySeller(t *testing.T) {
	deleteBySellerCommand, mocks := setUpDeleteProductsBySeller(t)

	idProd, _ := uuid.NewUUID()
	idSeller, _ := uuid.NewUUID()

	mocks.ProductRepository.EXPECT().GetFrom(idSeller).Return(nil, nil)
	mocks.ProductRepository.EXPECT().FindById(idProd).Return(&model.Product{}, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)

	err := deleteBySellerCommand.Execute(idSeller)

	assert.NoError(t, err)
}

func setUpDeleteProductsBySeller(t *testing.T) (*DeleteProductsBySeller, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	deleteProduct := NewDeleteProductEvent(mocks.ProductRepository)
	findProductsBySeller := query.NewFindProductsBySellerCommand(mocks.ProductRepository)

	return NewDeleteProductsBySellerEvent(findProductsBySeller, deleteProduct), mocks
}
