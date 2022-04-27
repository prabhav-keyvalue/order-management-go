// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/prabhav-keyvalue/order-management-go/entity"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// GetPriceByProductIds provides a mock function with given fields: productIds
func (_m *ProductRepository) GetPriceByProductIds(productIds []string) ([]entity.Product, error) {
	ret := _m.Called(productIds)

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func([]string) []entity.Product); ok {
		r0 = rf(productIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(productIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductRepository creates a new instance of ProductRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t testing.TB) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
