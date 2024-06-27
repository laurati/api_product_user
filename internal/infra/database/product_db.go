package database

import (
	"github.com/laurati/api_product_user/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) ProductInterface {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(user *entity.Product) error {
	return nil
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	return nil, nil
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	return nil, nil
}

func (p *Product) Update(product *entity.Product) error {
	return nil
}

func (p *Product) Delete(id string) error {
	return nil
}
