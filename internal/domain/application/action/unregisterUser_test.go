package action

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"testing"
)

func Test_UnregisterUser(t *testing.T) {
	unregisterUser, mocks := setUpUnregisterUser(t)

	id, _ := uuid.NewUUID()
	mocks.UserRepository.EXPECT().Delete(id).Return(nil)

	err := unregisterUser.Execute(id)

	assert.NoError(t, err)
}

func setUpUnregisterUser(t *testing.T) (*UnregisterUserEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)
	return NewUnregisterUserEvent(mocks.UserRepository), mocks
}
