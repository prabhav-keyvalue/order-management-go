package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	errorcode "github.com/prabhav-keyvalue/order-management-go/constant/error_code"
	"github.com/prabhav-keyvalue/order-management-go/dto"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/service"
	apiutil "github.com/prabhav-keyvalue/order-management-go/util/api"
	"gorm.io/gorm"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		orderService: service.NewOrderService(db),
	}
}

func (oc *OrderController) GetOrderById(c *gin.Context) {
	id := c.Param("id")
	order, err := oc.orderService.GetOrderById(id)

	if err != nil {
		apiutil.SendErrorResponse(c, http.StatusBadRequest, err.Error(), "order not found")
		return
	}

	apiutil.SendResponse(c, http.StatusOK, order)
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var createOrderInput dto.CreateOrderInputDto

	err := c.ShouldBindBodyWith(&createOrderInput, binding.JSON)

	if err != nil {
		logger.Error("Failed to create order", err.Error(), createOrderInput)
		apiutil.SendErrorResponse(c, http.StatusBadRequest, errorcode.INVALID_INPUT_CREATE_ORDER, "Failed to create order")
		return
	}

	order, err := oc.orderService.CreateOrder(createOrderInput)

	if err != nil {
		logger.Error("Failed to create order", createOrderInput)
		apiutil.SendErrorResponse(c, http.StatusBadRequest, err.Error(), "Failed to create order")
		return
	}

	apiutil.SendResponse(c, http.StatusOK, order)
}

func (oc *OrderController) EditOrder(c *gin.Context) {
	var editOrderInput dto.EditOrderInputDto

	err := c.ShouldBindBodyWith(&editOrderInput, binding.JSON)

	if err != nil {
		logger.Errorf("Failed to edit order | editOrderInput: %v", editOrderInput)
		apiutil.SendErrorResponse(c, http.StatusBadRequest, errorcode.INVALID_INPUT_EDIT_ORDER, "Failed to edit order")
		return
	}

	order, err := oc.orderService.EditOrder(editOrderInput)

	if err != nil {
		logger.Error("Failed to create order", editOrderInput)
		apiutil.SendErrorResponse(c, http.StatusBadRequest, errorcode.INVALID_INPUT_EDIT_ORDER, "Failed to edit order")
		return
	}

	apiutil.SendResponse(c, http.StatusOK, order)
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
	orderId := c.Param("id")

	_, err := oc.orderService.DeleteOrder(orderId)

	if err != nil {
		logger.Error("Failed to delete order", orderId)
		apiutil.SendErrorResponse(c, http.StatusBadRequest, errorcode.ORDER_DELETE_FAILED, "Failed to edit order")
		return
	}

	apiutil.SendResponse(c, http.StatusOK, orderId)
}
