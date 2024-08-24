package entity_test

import (
	"testing"

	"github.com/Gustavopnhro/api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "john@example.com"
	password := "securepassword123"

	user, err := entity.NewUser(entity.NewID().String(), name, email, password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.NotEqual(t, password, user.Password)
	assert.NotEmpty(t, user.ID)
	_, err = entity.ValidateID(user.ID.String())
	assert.NoError(t, err)
}

func TestValidatePassword(t *testing.T) {
	password := "securepassword123"
	user, err := entity.NewUser(entity.NewID().String(), "John Doe", "john@example.com", password)
	assert.NoError(t, err)

	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword("wrongpassword"))
}
