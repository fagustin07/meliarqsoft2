package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UpdateUserEvent(t *testing.T) {
	updateUser, mocks := setUpUpdateUser(t)

	id, _ := uuid.NewUUID()
	mocks.UserRepository.EXPECT().Update(id, "chester", "sandoval", "chester@gmail.com")

	assert.NoError(t, updateUser.Execute(id, "chester", "sandoval", "chester@gmail.com"))
}

func Test_UpdateUserEventOnInvalidEmailGivenThrowsError(t *testing.T) {
	updateUser, _ := setUpUpdateUser(t)
	id, _ := uuid.NewUUID()

	assert.Error(t, updateUser.Execute(id, "chester", "sandoval", "chester"))
}

func setUpUpdateUser(t *testing.T) (*UpdateUserEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewUpdateUserEvent(mocks.UserRepository), mocks
}
