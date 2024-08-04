package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("Jane Doo", "JaneDoo@jane.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "Jane Doo")
	assert.Equal(t, user.Email, "JaneDoo@jane.com")

}

func TestValidateUserPass(t *testing.T) {
	user, err := CreateUser("Jane Doo", "JaneDoo@jane.com", "1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "1234", user.Password)
}
