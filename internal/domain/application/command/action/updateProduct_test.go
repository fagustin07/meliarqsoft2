package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_UpdateProductEvent(t *testing.T) {
	updateProductEvent, mocks := setUpUpdateProduct(t)

	id, _ := uuid.NewUUID()
	prod, _ := dto.ProductDTO{
		ID:          id,
		Name:        "a",
		Description: "salado",
		Category:    "no",
		Price:       float32(20),
		Stock:       50,
	}.MapToModel()

	mocks.ProductRepository.EXPECT().Update(id, "f", "dulce", "ok", float32(43), 40).Return(prod, nil)

	prodGenerated, err := updateProductEvent.Execute(id, "f", "dulce", "ok", float32(43), 40)

	assert.NoError(t, err)
	assert.EqualValues(t, prod, prodGenerated)
}

func setUpUpdateProduct(t *testing.T) (*UpdateProductEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUpdateProductEvent(mocks.ProductRepository), mocks
}
