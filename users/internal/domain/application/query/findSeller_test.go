package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_FindSellerEvent(t *testing.T) {
	findSellerEvent, mocks := setUpFindSellerEvent(t)

	id, err := uuid.NewUUID()

	seller, _ := dto.SellerDTO{
		ID:           id,
		BusinessName: "Marolio",
		Email:        "marolio@ventas.com",
	}.MapToModel()
	sellersList := []*model.Seller{&seller}

	mocks.SellerRepository.EXPECT().Find("MAr").Return(sellersList, nil)

	sellersRes, err := findSellerEvent.Execute("MAr")

	assert.NoError(t, err)
	assert.Len(t, sellersRes, 1)
	assert.Equal(t, sellersRes[0].ID, seller.ID)
	assert.Equal(t, sellersRes[0].BusinessName, seller.BusinessName)
	assert.Equal(t, sellersRes[0].Email, seller.Email)

}

func setUpFindSellerEvent(t *testing.T) (*FindSellerEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindSellerEvent(mocks.SellerRepository), mocks
}
