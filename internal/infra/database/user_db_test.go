package database

import (
	"api-golang/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "j@j.oi", "123456")
	userDB := NewUser(db)
	_, err = userDB.CreateUser(user)
	assert.NoError(t, err)
	var userFound *entity.User
	err = db.First(&userFound, "email = ?", user.Email).Error
	assert.NoError(t, err)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "j@j.oi", "123456")
	userDB := NewUser(db)
	_, err = userDB.CreateUser(user)
	assert.NoError(t, err)

	userFound, err := userDB.FindUserByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
	assert.Equal(t, userFound.Password, user.Password)	
}