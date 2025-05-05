package database

import "api-golang/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update (product *entity.Product) (error)
	Delete(id string) error
}