package query

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"testing"
)

func Test_ExistUser(t *testing.T) {
	existUser, mocks := setUpExistUser(t)

	id, _ := uuid.NewUUID()
	mocks.UserRepository.EXPECT().FindById(id).Return(&model.User{}, nil)

	err := existUser.Execute(id)

	assert.NoError(t, err)
}

func Test_ExistUserThrowsErrorWhenNotFoundUser(t *testing.T) {
	existUser, mocks := setUpExistUser(t)

	id, _ := uuid.NewUUID()
	mocks.UserRepository.EXPECT().FindById(id).Return(&model.User{}, errors.New(""))

	err := existUser.Execute(id)

	assert.Error(t, err)
}

func setUpExistUser(t *testing.T) (*ExistUser, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewExistUserCommand(mocks.UserRepository), mocks
}
