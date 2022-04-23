package service

import (
	"errors"

	"github.com/google/uuid"
	errorcode "github.com/prabhav-keyvalue/order-management-go/constant/error_code"
	"github.com/prabhav-keyvalue/order-management-go/dto"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/repository"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrderById(id string) (entity.Order, error)
	CreateOrder(createOrderInput dto.CreateOrderInputDto) (entity.Order, error)
	EditOrder(editOrderInput dto.EditOrderInputDto) (entity.Order, error)
	DeleteOrder(orderId string) (string, error)
}

type OrderServiceImpl struct {
	dB               *gorm.DB
	orderRepository  repository.OrderRepository
	orderItemService OrderItemService
	productService   ProductService
}

func NewOrderService(db *gorm.DB) OrderService {
	return &OrderServiceImpl{
		dB:               db,
		orderRepository:  repository.NewOrderRepository(db),
		orderItemService: NewOrderItemService(db),
		productService:   NewProductService(db),
	}
}

func (os *OrderServiceImpl) GetOrderById(id string) (order entity.Order, err error) {
	order, err = os.orderRepository.GetOrderById(id)

	if err != nil {
		logger.Errorf("Error geting order by id | Error: %v\n", id, err.Error())
		return order, errors.New(errorcode.ORDER_NOT_FOUND)
	}

	return order, err

}

func (os *OrderServiceImpl) CreateOrder(createOrderInput dto.CreateOrderInputDto) (order entity.Order, err error) {
	var newOrder entity.Order

	totalQuantity := 0

	for _, oi := range createOrderInput.OrderItems {
		totalQuantity += oi.Quantity
	}

	var productIds []string

	for _, oi := range createOrderInput.OrderItems {
		productIds = append(productIds, oi.ProductId)
	}

	productIdPrice, err := os.productService.GetPriceByProductIds(productIds)

	if err != nil {
		return newOrder, err
	}

	productIdPriceMap := make(map[string]float64)

	for _, p := range productIdPrice {
		productIdPriceMap[p.Id] = p.UnitPrice
	}

	var totalPrice float64 = 0

	for _, oi := range createOrderInput.OrderItems {
		totalPrice += productIdPriceMap[oi.ProductId] * float64(oi.Quantity)
	}

	newOrder = entity.Order{
		CustomerId:    createOrderInput.CustomerId,
		TotalQuantity: totalQuantity,
		TotalPrice:    totalPrice,
	}

	newOrder.Id = uuid.NewString()

	tx := os.dB.Begin()
	order, err = os.orderRepository.CreateOrder(newOrder, tx)

	if err != nil {
		logger.Errorf("Failed to insert into order table, createOrderInput: %v | Error: %v\n", createOrderInput, err.Error())
		tx.Rollback()
		return order, err
	}

	logger.Infof("Created Order %v", order)

	var createOrderItemsInput []entity.OrderItem

	productQuantityMap := make(map[string]int)

	for _, p := range createOrderInput.OrderItems {
		productQuantityMap[p.ProductId] += p.Quantity
	}

	for id, q := range productQuantityMap {
		createOrderItemsInput = append(createOrderItemsInput, entity.OrderItem{
			OrderId:   order.Id,
			ProductId: id,
			Quantity:  q,
			Price:     productIdPriceMap[id],
			RowTotal:  productIdPriceMap[id] * float64(q),
		})
	}

	orderItems, err := os.orderItemService.CreateOrderItemsWithOrder(createOrderItemsInput, tx)

	if err != nil {
		tx.Rollback()
		return order, err
	}

	logger.Info("Created order items", orderItems)

	tx.Commit()

	return order, err
}

func (os *OrderServiceImpl) EditOrder(editOrderInput dto.EditOrderInputDto) (entity.Order, error) {
	totalQuantity := 0
	var newOrder entity.Order
	for _, oi := range editOrderInput.OrderItems {
		totalQuantity += oi.Quantity
	}

	existingOrder, err := os.GetOrderById(editOrderInput.OrderId)

	if err != nil {
		logger.Errorf("Could not find order | orderId: %v | Error: %v", editOrderInput.OrderId, err.Error())
		return newOrder, err
	}

	if existingOrder.CustomerId != editOrderInput.CustomerId {
		logger.Errorf("Order does not belong to customer | orderId: %v | Error: %v", editOrderInput.OrderId, err.Error())
		return newOrder, err
	}

	var productIds []string

	for _, oi := range editOrderInput.OrderItems {
		productIds = append(productIds, oi.ProductId)
	}

	productIdPrice, err := os.productService.GetPriceByProductIds(productIds)

	if err != nil {
		return newOrder, err
	}

	productIdPriceMap := make(map[string]float64)

	for _, p := range productIdPrice {
		productIdPriceMap[p.Id] = p.UnitPrice
	}

	var totalPrice float64 = 0

	for _, oi := range editOrderInput.OrderItems {
		totalPrice += productIdPriceMap[oi.ProductId] * float64(oi.Quantity)
	}

	tx := os.dB.Begin()

	if totalPrice != existingOrder.TotalPrice || totalQuantity != existingOrder.TotalQuantity {
		existingOrder.TotalPrice = totalPrice
		existingOrder.TotalQuantity = totalQuantity
		newOrder, err = os.orderRepository.EditOrder(existingOrder, tx)

		if err != nil {
			logger.Errorf("Failed to update order | editOrderInput: %v | Error: %v\n", editOrderInput, err.Error())
			tx.Rollback()
			return newOrder, err
		}

	}

	// create new order items if any

	existingOrderItems, err := os.orderItemService.GetOrderItemsByOrderId(editOrderInput.OrderId, tx)

	if err != nil {
		tx.Rollback()
		return newOrder, err
	}

	var newOrderItems []entity.OrderItem

	existingOrderItemProductIdQuantityMap := make(map[string]int)

	for _, oit := range existingOrderItems {
		existingOrderItemProductIdQuantityMap[oit.ProductId] = oit.Quantity
	}

	for _, eoi := range editOrderInput.OrderItems {
		if _, ok := existingOrderItemProductIdQuantityMap[eoi.ProductId]; !ok {
			newOrderItems = append(newOrderItems, entity.OrderItem{
				OrderId:   editOrderInput.OrderId,
				ProductId: eoi.ProductId,
				Quantity:  eoi.Quantity,
				Price:     productIdPriceMap[eoi.ProductId],
				RowTotal:  productIdPriceMap[eoi.ProductId] * float64(eoi.Quantity),
			})
		}
	}

	if len(newOrderItems) > 0 {
		newOrderItems, err = os.orderItemService.CreateOrderItemsWithOrder(newOrderItems, tx)

		if err != nil {
			tx.Rollback()
			return newOrder, err
		}

		logger.Infof("Created new order items: %v\n", newOrderItems)
	}

	// deleting order items

	var orderItemsToBeDeleted []string

	for _, oit := range existingOrderItems {
		if _, ok := productIdPriceMap[oit.ProductId]; !ok {
			orderItemsToBeDeleted = append(orderItemsToBeDeleted, oit.Id)
		}
	}

	if len(orderItemsToBeDeleted) > 0 {
		err := os.orderItemService.DeleteOrderItemsByIds(orderItemsToBeDeleted, tx)

		if err != nil {
			tx.Rollback()
			return newOrder, err
		}
	}

	// changing quantity of existing order

	var updatedOrderItems []entity.OrderItem

	for _, oi := range editOrderInput.OrderItems {
		if val, ok := existingOrderItemProductIdQuantityMap[oi.ProductId]; ok && oi.Quantity != val {
			updatedOrderItems = append(updatedOrderItems, entity.OrderItem{
				OrderId:   editOrderInput.OrderId,
				ProductId: oi.ProductId,
				Quantity:  oi.Quantity,
				Price:     productIdPriceMap[oi.ProductId],
				RowTotal:  productIdPriceMap[oi.ProductId] * float64(oi.Quantity),
			})
		}
	}

	for _, oi := range updatedOrderItems {
		err = os.orderItemService.UpdateOrderItemQuantity(oi, tx)

		if err != nil {
			tx.Rollback()
			return newOrder, err
		}
	}

	tx.Commit()

	return newOrder, err

}

func (os *OrderServiceImpl) DeleteOrder(orderId string) (string, error) {
	tx := os.dB.Begin()

	err := os.orderRepository.DeleteOrder(orderId)

	if err != nil {
		logger.Errorf("Delete order failed | orderId: %v | Error: %v", orderId, err.Error())
		tx.Rollback()
		return orderId, err
	}

	_, err = os.orderItemService.DeleteOrderItemsByOrderId(orderId)

	if err != nil {
		tx.Rollback()
		return orderId, err
	}
	tx.Commit()

	return orderId, err
}
