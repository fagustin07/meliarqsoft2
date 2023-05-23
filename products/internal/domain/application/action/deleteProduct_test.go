package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_DeleteProductEvent(t *testing.T) {
	deleteProdEvent, mocks := setUpDeleteProdEvent(t)

	idProd, _ := uuid.NewUUID()

	mocks.ProductRepository.EXPECT().FindById(idProd).Return(nil, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.DeletePurchases.EXPECT().Execute([]uuid.UUID{idProd}).Return(nil)

	err := deleteProdEvent.Execute(idProd)

	assert.NoError(t, err)
}

func Test_DeleteProductEventFailsWhenDeletePurchasesFailsAndRestoreProduct(t *testing.T) {
	deleteProdEvent, mocks := setUpDeleteProdEvent(t)

	idProd, _ := uuid.NewUUID()

	mocks.ProductRepository.EXPECT().FindById(idProd).Return(nil, nil)
	mocks.ProductRepository.EXPECT().Delete(idProd).Return(nil)
	mocks.DeletePurchases.EXPECT().Execute([]uuid.UUID{idProd}).Return(errors.New(""))
	mocks.ProductRepository.EXPECT().Restore([]uuid.UUID{idProd}).Return(int64(1), nil)

	err := deleteProdEvent.Execute(idProd)

	assert.Error(t, err)
}

func setUpDeleteProdEvent(t *testing.T) (*DeleteProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewDeleteProductEvent(mocks.ProductRepository, mocks.DeletePurchases), mocks
}
