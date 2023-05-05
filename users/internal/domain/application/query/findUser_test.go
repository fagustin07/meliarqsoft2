package query

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meliarqsoft2/internal/domain/mock"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"testing"
)

func Test_FindUserEvent(t *testing.T) {
	findUserEvent, mocks := setUpFindUserEvent(t)

	id, err := uuid.NewUUID()

	user, _ := dto.UserDTO{
		ID:      id,
		Surname: "sanchez",
		Name:    "chester",
		Email:   "chester99@gmail.com",
	}.MapToModel()
	usersList := []*model.User{&user}

	mocks.UserRepository.EXPECT().Find("99").Return(usersList, nil)

	usersRes, err := findUserEvent.Execute("99")

	assert.NoError(t, err)
	assert.Len(t, usersRes, 1)
	assert.Equal(t, usersRes[0].ID, user.ID)
	assert.Equal(t, usersRes[0].Name, user.Name)
	assert.Equal(t, usersRes[0].Surname, user.Surname)
	assert.Equal(t, usersRes[0].Email, user.Email)

}

func setUpFindUserEvent(t *testing.T) (*FindUserEvent, *mock.RepositoriesMock) {
	mocks := mock.NewMockRepositories(t)

	return NewFindUserEvent(mocks.UserRepository), mocks
}
