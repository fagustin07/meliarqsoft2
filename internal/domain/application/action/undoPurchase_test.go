package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UndoPurchase(t *testing.T) {
	undoPurchase, mocks := setUpUndoPurchase(t)

	id, _ := uuid.NewUUID()
	mocks.PurchaseRepository.EXPECT().Delete(id).Return(nil)

	assert.NoError(t, undoPurchase.Execute(id))
}

func setUpUndoPurchase(t *testing.T) (*UndoPurchase, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUndoPurchaseCommand(mocks.PurchaseRepository), mocks
}
