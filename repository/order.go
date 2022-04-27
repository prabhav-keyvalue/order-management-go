package repository

import (
	"fmt"
	"strconv"

	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/dto"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderById(id string) (entity.Order, error)
	CreateOrder(orderInput entity.Order, tx *gorm.DB) (entity.Order, error)
	EditOrder(editOrderInput entity.Order, tx *gorm.DB) (entity.Order, error)
	DeleteOrder(orderId string, tx ...*gorm.DB) error
	GetOrders(orderFilterParams dto.OrderFilterParams, paginationInput dto.PaginationParams, sortOptions dto.SortOptions) ([]entity.Order, model.PageInfo, error)
}

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func applyOrderFilters(orderFilterParams dto.OrderFilterParams, qb *gorm.DB) *gorm.DB {
	if f := orderFilterParams.MaxQuantity; f != "" {
		if val, err := strconv.Atoi(f); err == nil {
			qb = qb.Where("total_quantity <= ?", val)
		}
	}

	if f := orderFilterParams.MinQuantity; f != "" {
		if val, err := strconv.Atoi(f); err == nil {
			qb = qb.Where("total_quantity >= ?", val)
		}
	}

	return qb
}

func (or *OrderRepositoryImpl) GetOrders(orderFilterParams dto.OrderFilterParams, paginationInput dto.PaginationParams, sortOptions dto.SortOptions) ([]entity.Order, model.PageInfo, error) {
	var orders []entity.Order
	var count int64
	qb := applyOrderFilters(orderFilterParams, or.DB)

	qb = qb.Scopes(Paginate(paginationInput.Limit, paginationInput.Offset))
	qb = qb.Order(fmt.Sprintf("%s %s", sortOptions.SortKey, sortOptions.SortOrder))

	err := qb.Find(&orders).Offset(-1).Limit(-1).Count(&count).Error

	pageInfo := model.PageInfo{
		TotalCount: count,
		Offset:     paginationInput.Offset,
		Limit:      paginationInput.Limit,
	}
	return orders, pageInfo, err
}

func (or *OrderRepositoryImpl) GetOrderById(id string) (entity.Order, error) {
	var order entity.Order
	err := or.DB.Preload("OrderItems").Preload("Customer").Find(&order, "id = ?", id).Error

	return order, err
}

func (or *OrderRepositoryImpl) CreateOrder(orderInput entity.Order, tx *gorm.DB) (entity.Order, error) {
	newOrder := orderInput

	err := tx.Create(&newOrder).Error

	return newOrder, err
}

func (or *OrderRepositoryImpl) EditOrder(editOrderInput entity.Order, tx *gorm.DB) (entity.Order, error) {
	err := tx.Table(db.GetTableNameWithSchema("order")).Where("id = ?", editOrderInput.Id).Updates(entity.Order{TotalQuantity: editOrderInput.TotalQuantity, TotalPrice: editOrderInput.TotalPrice}).Error

	return editOrderInput, err
}

func (or *OrderRepositoryImpl) DeleteOrder(orderId string, tx ...*gorm.DB) error {
	var db *gorm.DB

	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = or.DB
	}
	return db.Delete(&entity.Order{}, "id = ?", orderId).Error
}
