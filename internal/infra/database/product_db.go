package database

import (
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"gorm.io/gorm"
)

type ProductDBInterface interface {
	Create(product *entity.Product) error
	FindAll(page int, limit int, sort string) (*[]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) FindAll(page int, limit int, sort string) (*[]entity.Product, error) {
	var products []entity.Product
	if err := p.DB.Offset((page) * limit).Limit(limit).Order("created_at " + sort).Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductDB) Delete(id string) error {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}
	return p.DB.Delete(&product).Error
}
