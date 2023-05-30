package store_test

import (
	"github.com/greatvill/survey.git/internal/app/model"
	"github.com/greatvill/survey.git/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@email.ru",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	_, err = s.User().Create(&model.User{
		Email: "user@email.ru",
	})

	var u *model.User
	u, err = s.User().FindByEmail("user@email.ru")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
