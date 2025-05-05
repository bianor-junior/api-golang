package database

import (
	"api-golang/internal/entity"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (u *Product) Create(product *entity.Product) error {
	return u.DB.Create(product).Error
}

func (u *Product) FindById(id string) (*entity.Product, error) {
	var existingProduct entity.Product
	if err := u.DB.First(&existingProduct, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &existingProduct, nil
}

func (u *Product) Update(product *entity.Product) (error) {
	var existingProduct entity.Product
	err := u.DB.First(&existingProduct, "id = ?", product.ID).Error
	if err != nil {
		return err
	}
	return u.DB.Save(product).Error
}

func (u *Product) Delete(id string) error {
	var existingProduct entity.Product
	if err := u.DB.First(&existingProduct, "id = ?", id).Error; err != nil {
		return err
	}
	return u.DB.Delete(&entity.Product{}, id).Error
}

func (u *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var product []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&product).Error
	} else {
		err = u.DB.Order("created_at " + sort).Find(&product).Error
	}
	return product, err
}