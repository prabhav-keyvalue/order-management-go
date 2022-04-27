package service

import (
	"github.com/google/uuid"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/repository"
	"gorm.io/gorm"
)

type OrderItemService interface {
	CreateOrderItemsWithOrder(createOrderItemsInput []entity.OrderItem, tx *gorm.DB) ([]entity.OrderItem, error)
	GetOrderItemsByOrderId(orderId string, tx *gorm.DB) ([]entity.OrderItem, error)
	DeleteOrderItemsByIds(orderItemIds []string, tx *gorm.DB) error
	UpdateOrderItemQuantity(orderItem entity.OrderItem, tx *gorm.DB) error
	DeleteOrderItemsByOrderId(orderId string, tx ...*gorm.DB) error
}

type OrderItemServiceImpl struct {
	orderItemRepository repository.OrderItemRepository
}

func NewOrderItemService(db *gorm.DB) OrderItemService {
	return &OrderItemServiceImpl{
		orderItemRepository: repository.NewOrderItemRepository(db),
	}
}

func (oit *OrderItemServiceImpl) CreateOrderItemsWithOrder(createOrderItemsInput []entity.OrderItem, tx *gorm.DB) ([]entity.OrderItem, error) {
	var newOrderItems []entity.OrderItem

	for _, o := range createOrderItemsInput {
		newOrderItem := entity.OrderItem{
			OrderId:   o.OrderId,
			ProductId: o.ProductId,
			Price:     o.Price,
			Quantity:  o.Quantity,
			RowTotal:  o.Price * float64(o.Quantity),
		}
		newOrderItem.Id = uuid.NewString()
		newOrderItems = append(newOrderItems, newOrderItem)
	}

	orderItems, err := oit.orderItemRepository.CreateOrderItemsWithOrder(newOrderItems, tx)

	if err != nil {
		logger.Errorf("Failed to create order items | Error: %v\n", err.Error())
	}

	return orderItems, err
}

func (oit *OrderItemServiceImpl) GetOrderItemsByOrderId(orderId string, tx *gorm.DB) ([]entity.OrderItem, error) {
	orderItems, err := oit.orderItemRepository.GetOrderItemsByOrderId(orderId, tx)

	if err != nil {
		logger.Errorf("Failed to fetch order items | orderId %v | Error: %v\n", orderId, err.Error())
	}

	return orderItems, err
}

func (oit *OrderItemServiceImpl) DeleteOrderItemsByIds(orderItemIds []string, tx *gorm.DB) error {
	err := oit.orderItemRepository.DeleteOrderItemsByIds(orderItemIds, tx)

	if err != nil {
		logger.Errorf("Failed to delete order items | orderItemsIds %v", orderItemIds)
	}

	return err
}

func (oit *OrderItemServiceImpl) UpdateOrderItemQuantity(orderItem entity.OrderItem, tx *gorm.DB) error {
	err := oit.orderItemRepository.UpdateOrderItemQuantity(orderItem, tx)

	if err != nil {
		logger.Error("Failed to update order quantity | orderItem: %v | Error: %v\n", orderItem, err.Error())
	}

	return err
}

func (oit *OrderItemServiceImpl) DeleteOrderItemsByOrderId(orderId string, tx ...*gorm.DB) (err error) {

	err = oit.orderItemRepository.DeleteOrderItemsByOrderId(orderId, tx...)

	if err != nil {
		logger.Errorf("Failed to delete order item | orderId: %v | Error: %v\n", orderId, err.Error())
	}
	return
}
