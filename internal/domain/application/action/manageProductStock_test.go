package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_ManageProductStock_ConsumeStock(t *testing.T) {
	manageProductStock, mocks := setUpManageProductStock(t)

	id, _ := uuid.NewUUID()
	prod, _ := dto.ProductDTO{
		ID:          id,
		Name:        "a",
		Description: "salado",
		Category:    "no",
		Price:       float32(20),
		Stock:       50,
	}.MapToModel()

	unitsToConsume := 10
	stockAfterConsume := prod.Stock.Amount - unitsToConsume
	mocks.ProductRepository.EXPECT().FindById(id).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().UpdateStock(id, stockAfterConsume).Return(nil)

	err := manageProductStock.Execute(id, unitsToConsume, false)

	assert.NoError(t, err)
	assert.Equal(t, prod.Stock.Amount, stockAfterConsume)
}

func Test_ManageProductStock_RestoreStock(t *testing.T) {
	manageProductStock, mocks := setUpManageProductStock(t)

	id, _ := uuid.NewUUID()
	prod, _ := dto.ProductDTO{
		ID:          id,
		Name:        "a",
		Description: "salado",
		Category:    "no",
		Price:       float32(20),
		Stock:       50,
	}.MapToModel()

	unitsToRestore := 10
	stockAfterRestore := prod.Stock.Amount + unitsToRestore
	mocks.ProductRepository.EXPECT().FindById(id).Return(&prod, nil)
	mocks.ProductRepository.EXPECT().UpdateStock(id, stockAfterRestore).Return(nil)

	err := manageProductStock.Execute(id, unitsToRestore, true)

	assert.NoError(t, err)
	assert.Equal(t, prod.Stock.Amount, stockAfterRestore)
}

func setUpManageProductStock(t *testing.T) (*ManageProductStock, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewManageProductStockCommand(mocks.ProductRepository), mocks
}
