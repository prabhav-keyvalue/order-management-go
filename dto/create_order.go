package dto

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	errorcode "github.com/prabhav-keyvalue/order-management-go/constant/error_code"
	"github.com/prabhav-keyvalue/order-management-go/logger"
)

type CreateOrderInputDto struct {
	*Dto
	CustomerId string      `json:"customerId"`
	OrderItems []OrderItem `json:"orderItems"`
}

type OrderItem struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

func NewCreateOrderInputDto() *CreateOrderInputDto {
	d := &Dto{}
	c := &CreateOrderInputDto{Dto: d}
	d.ISDto = c
	return c
}

func (ci *CreateOrderInputDto) validateDto(c *gin.Context) (err error) {
	var createOrderInput CreateOrderInputDto
	c.ShouldBindBodyWith(&createOrderInput, binding.JSON)
	err = validation.ValidateStruct(&createOrderInput,
		validation.Field(&createOrderInput.CustomerId, validation.Required, is.UUIDv4),
	)

	if err != nil {
		logger.Error("Invalid create order request", err.Error(), createOrderInput)
		return errors.New(errorcode.INVALID_INPUT_CREATE_ORDER)
	}

	productIdMap := make(map[string]struct{})

	for _, t := range createOrderInput.OrderItems {
		if _, ok := productIdMap[t.ProductId]; ok {
			logger.Error("Invalid create order request duplicate product id")
			return errors.New(errorcode.INVALID_INPUT_CREATE_ORDER)
		}
		err = validation.ValidateStruct(&t,
			validation.Field(&t.ProductId, validation.Required, is.UUIDv4),
			validation.Field(&t.Quantity, validation.Required, validation.Min(0)),
		)

		if err != nil {
			logger.Error("Invalid create order request", err.Error(), createOrderInput)
			return errors.New(errorcode.INVALID_INPUT_CREATE_ORDER)
		}

		productIdMap[t.ProductId] = struct{}{}
	}

	return

}
