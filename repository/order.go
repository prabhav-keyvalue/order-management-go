package repository

import (
	"fmt"

	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderById(id string) (entity.Order, error)
	CreateOrder(orderInput entity.Order, tx *gorm.DB) (entity.Order, error)
	EditOrder(editOrderInput entity.Order, tx *gorm.DB) (entity.Order, error)
	DeleteOrder(orderId string) error
}

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (or *OrderRepositoryImpl) GetOrderById(id string) (entity.Order, error) {
	var order entity.Order

	err := or.DB.Joins("inner join order_item on order_item.order_id = order.id").Where("order.id = id", id).Find(&order).Error

	// var orderitems []entity.OrderItem
	// err := or.DB.Table(db.GetTableNameWithSchema("order")).Where("id = ?", id).Scan(&order).Error
	// err := or.DB.Find(&order, "id = ?", id).Error
	// orderItemTableName := db.GetTableNameWithSchema("order_item")
	// err := or.DB.Model(&order).Select("*").Where().Association("OrderItems").Find(&orderitems)
	// err := or.DB.Table("order").Select("*").Where("id = ?", id).Scan(&order).Error
	// err := or.DB.Raw("SELECT * FROM test.order o inner join test.order_item oi on oi.order_id = o.id where o.id = ? and o.deleted_at is null", id).Find(&order).Error
	fmt.Println("sdfsdf", order)
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

func (or *OrderRepositoryImpl) DeleteOrder(orderId string) error {
	return or.DB.Table(db.GetTableNameWithSchema("order")).Where("id = ?", orderId).Delete(&entity.Order{}).Error
}
