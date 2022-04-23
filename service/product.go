package service

import (
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	GetPriceByProductIds(productIds []string) ([]entity.Product, error)
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		productRepository: repository.NewProductRepository(db),
	}
}

func (ps *ProductServiceImpl) GetPriceByProductIds(productIds []string) ([]entity.Product, error) {
	products, err := ps.productRepository.GetPriceByProductIds(productIds)

	if err != nil {
		logger.Errorf("Failed to get product prices | Error: %v | productIds: %v", err.Error(), productIds)
	}

	return products, err
}
