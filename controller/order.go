package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	errorcode "github.com/prabhav-keyvalue/order-management-go/constant/error_code"
	"github.com/prabhav-keyvalue/order-management-go/dto"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/model"
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

// GetOrderById godoc
// @ID get-order-by-id
// @Summary To get a single Order
// @Description to get a single Order ( Fetch All Information related to order)
// @Tags orders
// @Param        id   path      string  true  "order id"
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=entity.Order,error=model.ApiError{code=string,message=string}}
// @Router /orders/{id} [get]
func (oc *OrderController) GetOrderById(c *gin.Context) {
	id := c.Param("id")
	order, err := oc.orderService.GetOrderById(id)

	if err != nil {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    err.Error(),
				Message: "get order failed",
			},
		})
		return
	}

	apiutil.SendResponse(c, http.StatusOK, &model.Response{
		Data: order,
	})
}

// CreateOrder godoc
// @ID create-order
// @Summary Create an order
// @Description Create an Order
// @Tags orders
// @Param        CreateOrderInputDto   body      dto.CreateOrderInputDto  true  "CreateOrderInputDto"
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=entity.Order,error=model.ApiError{code=string,message=string}}
// @Router /orders [post]
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var createOrderInput dto.CreateOrderInputDto

	err := c.ShouldBindBodyWith(&createOrderInput, binding.JSON)

	if err != nil {
		logger.Error("Failed to create order", err.Error(), createOrderInput)
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INVALID_INPUT_CREATE_ORDER,
				Message: "Failed to create order",
			},
		})
		return
	}

	order, err := oc.orderService.CreateOrder(createOrderInput)

	if err != nil {
		logger.Error("Failed to create order", createOrderInput)
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to create order",
			},
		})
		return
	}

	apiutil.SendResponse(c, http.StatusOK, &model.Response{
		Data: order,
	})
}

func (oc *OrderController) EditOrder(c *gin.Context) {
	var editOrderInput dto.EditOrderInputDto

	err := c.ShouldBindBodyWith(&editOrderInput, binding.JSON)

	if err != nil {
		logger.Errorf("Failed to edit order | editOrderInput: %v", editOrderInput)
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INVALID_INPUT_EDIT_ORDER,
				Message: "Failed to edit order",
			},
		})
		return
	}

	order, err := oc.orderService.EditOrder(editOrderInput)

	if err != nil {
		logger.Error("Failed to edit order", editOrderInput)
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to edit order",
			},
		})
		return
	}

	apiutil.SendResponse(c, http.StatusOK, &model.Response{
		Data: order,
	})
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
	orderId := c.Param("id")

	err := oc.orderService.DeleteOrder(orderId)

	if err != nil {
		logger.Error("Failed to delete order", orderId)
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.ORDER_DELETE_FAILED,
				Message: "Failed to delete order",
			},
		})
		return
	}

	apiutil.SendResponse(c, http.StatusOK, &model.Response{
		Data: orderId,
	})
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	val, exists := c.Get("orderFilterParams")
	if !exists {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	orderFilterParams, ok := val.(dto.OrderFilterParams)

	if !ok {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	val, exists = c.Get("PaginationQueryParams")
	if !exists {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	paginationInput, ok := val.(dto.PaginationParams)

	if !ok {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	val, exists = c.Get("SortOptionsQueryParams")
	if !exists {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	sortOptions, ok := val.(dto.SortOptions)

	if !ok {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	orders, pageInfo, err := oc.orderService.GetOrders(orderFilterParams, paginationInput, sortOptions)

	if err != nil {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code:    errorcode.INTERNAL_SERVER_ERROR,
				Message: "Failed to get orders",
			},
		})
		return
	}

	apiutil.SendResponse(c, http.StatusOK, &model.Response{
		Data:     orders,
		PageInfo: &pageInfo,
	})
}
