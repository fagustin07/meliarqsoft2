package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UpdateSellerEvent(t *testing.T) {
	updateSeller, mocks := setUpUpdateSeller(t)

	id, _ := uuid.NewUUID()
	mocks.SellerRepository.EXPECT().Update(id, "marolio", "marolio@ventas.com")

	assert.NoError(t, updateSeller.Execute(id, "marolio", "marolio@ventas.com"))
}

func Test_UpdateSellerEventOnInvalidEmailGivenThrowsError(t *testing.T) {
	updateSeller, _ := setUpUpdateSeller(t)
	id, _ := uuid.NewUUID()

	assert.Error(t, updateSeller.Execute(id, "marolio", "marolio@.com"))
}

func setUpUpdateSeller(t *testing.T) (*UpdateSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUpdateSellerEvent(mocks.SellerRepository), mocks
}
