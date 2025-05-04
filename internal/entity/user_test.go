package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// Test the NewUser function
	user, err := NewUser("John Doe", "j@j.oi", "1234567890")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.oi", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.oi", "1234567890")
	assert.NoError(t, err)
	assert.True(t, user.ValidatePassword("1234567890"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, user.Password, "1234567890")
}
