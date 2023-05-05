package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_RegisterUserEvent(t *testing.T) {
	registerUserEvent, mocks := setUpRegisterUser(t)

	id, _ := uuid.NewUUID()
	user := &model.User{}

	mocks.UserRepository.EXPECT().Create(user).Return(id, nil)

	newId, err := registerUserEvent.Execute(user)

	assert.Equal(t, id, newId)
	assert.NoError(t, err)
}

func Test_RegisterUserEvent_ThrowsError(t *testing.T) {
	registerUserEvent, mocks := setUpRegisterUser(t)

	user := &model.User{}

	mocks.UserRepository.EXPECT().Create(user).Return(uuid.Nil, errors.New(""))

	blankID, err := registerUserEvent.Execute(user)

	assert.Equal(t, blankID, uuid.Nil)
	assert.Error(t, err)
}

func setUpRegisterUser(t *testing.T) (*RegisterUserEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewRegisterUserEvent(mocks.UserRepository), mocks
}
