package user

import (
	"testing"
	"time"

	"github.com/AhmetSBulbul/golang-api-playground/entity"
	"github.com/stretchr/testify/assert"
)

// Test Cases
// - valid search
// - invalid search
// - valid list users
// - valid get users
// - test update
// - test delete

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		Email:     "mail@test.app",
		Password:  "Password123",
		FirstName: "Jack",
		LastName:  "Sparrow",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.FirstName = "Lara"
	uID, _ := m.CreateUser(u1.Email, u1.Password, u1.FirstName, u1.LastName)
	_, _ = m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)
	t.Run("search", func(t *testing.T) {
		c, err := m.SearchUsers("jack")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Sparrow", c[0].LastName)

		c, err = m.SearchUsers("gina")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("listAll", func(t *testing.T) {
		all, err := m.ListUsers()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetUser(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.FirstName, saved.FirstName)
	})
}
