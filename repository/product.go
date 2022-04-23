package repository

import (
	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetPriceByProductIds(productIds []string) ([]entity.Product, error)
}

type ProductRepositoryImpl struct {
	dB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		dB: db,
	}
}

func (pr *ProductRepositoryImpl) GetPriceByProductIds(productIds []string) ([]entity.Product, error) {
	var products []entity.Product
	err := pr.dB.Table(db.GetTableNameWithSchema("product")).Select("unit_price, id").Where("id in ?", productIds).Find(&products).Error
	return products, err
}
