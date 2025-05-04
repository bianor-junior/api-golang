package entity

import (
	"api-golang/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIDIsRequired = errors.New("ID is required")
	ErrIvalidID = errors.New("ID is invalid")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrIvalidPrice = errors.New("price is invalid")

)

type Product struct {
	ID   entity.ID
	Name string
	Price int
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrIvalidPrice
	}
	return nil
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}