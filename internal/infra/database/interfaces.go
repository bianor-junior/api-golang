package database

import "api-golang/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}