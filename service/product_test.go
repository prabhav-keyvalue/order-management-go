package service

import (
	"testing"

	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func GetNewProductService(repo *mocks.ProductRepository) *ProductServiceImpl {

	service := ProductServiceImpl{productRepository: repo}

	return &service
}

func GetListOfProducts() []entity.Product {
	return []entity.Product{
		{
			BaseEntity: entity.BaseEntity{Id: "id"},
		},
	}
}

func TestGetPriceByProductIds(t *testing.T) {

	t.Run("should return list of productids successfully", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		productIds := []string{"id1", "id2"}
		products := GetListOfProducts()
		repo.On("GetPriceByProductIds", productIds).Return(products, nil)

		service := GetNewProductService(repo)

		res, err := service.GetPriceByProductIds(productIds)

		assert.Equal(t, products, res)
		assert.Nil(t, err)
	})

}
