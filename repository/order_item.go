package repository

import (
	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	CreateOrderItemsWithOrder(orderItems []entity.OrderItem, tx *gorm.DB) ([]entity.OrderItem, error)
	GetOrderItemsByOrderId(orderId string, tx *gorm.DB) ([]entity.OrderItem, error)
	DeleteOrderItemsByIds(orderItemIds []string, tx *gorm.DB) error
	UpdateOrderItemQuantity(orderItem entity.OrderItem, tx *gorm.DB) error
	DeleteOrderItemsByOrderId(orderId string, tx ...*gorm.DB) error
}

type OrderItemRepositoryImpl struct {
	dB *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepositoryImpl {
	return &OrderItemRepositoryImpl{
		dB: db,
	}
}

func (oit *OrderItemRepositoryImpl) CreateOrderItemsWithOrder(orderItems []entity.OrderItem, tx *gorm.DB) ([]entity.OrderItem, error) {
	newOrderItems := orderItems

	err := tx.Create(&newOrderItems).Error

	return newOrderItems, err
}

func (oit *OrderItemRepositoryImpl) GetOrderItemsByOrderId(orderId string, tx *gorm.DB) ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	err := tx.Table(db.GetTableNameWithSchema("order_item")).Select("*").Where("order_id = ?", orderId).Find(&orderItems).Error
	return orderItems, err
}

func (oit *OrderItemRepositoryImpl) DeleteOrderItemsByIds(orderItemIds []string, tx *gorm.DB) error {
	err := tx.Table(db.GetTableNameWithSchema("order_item")).Where("id IN ?", orderItemIds).Delete(&entity.OrderItem{}).Error

	return err
}

func (oit *OrderItemRepositoryImpl) DeleteOrderItemsByOrderId(orderId string, tx ...*gorm.DB) error {
	var db *gorm.DB

	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = oit.dB
	}

	return db.Delete(&entity.OrderItem{}, "order_id = ?", orderId).Error
}

func (oit *OrderItemRepositoryImpl) UpdateOrderItemQuantity(orderItem entity.OrderItem, tx *gorm.DB) error {
	err := tx.Table(db.GetTableNameWithSchema("order_item")).Where("order_id = ? AND product_id = ? ", orderItem.OrderId, orderItem.ProductId).Updates(&entity.OrderItem{
		Quantity: orderItem.Quantity,
		Price:    orderItem.Price,
		RowTotal: orderItem.RowTotal,
	}).Error

	return err
}
