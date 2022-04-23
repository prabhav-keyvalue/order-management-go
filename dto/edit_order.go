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

type EditOrderInputDto struct {
	*Dto
	OrderId    string      `json:"orderId"`
	CustomerId string      `json:"customerId"`
	OrderItems []OrderItem `json:"orderItems"`
}

func NewEditOrderInputDto() *EditOrderInputDto {
	d := &Dto{}
	e := &EditOrderInputDto{Dto: d}
	d.ISDto = e
	return e
}

func (ei *EditOrderInputDto) validateDto(c *gin.Context) (err error) {
	var editOrderInput EditOrderInputDto
	c.ShouldBindBodyWith(&editOrderInput, binding.JSON)
	err = validation.ValidateStruct(&editOrderInput,
		validation.Field(&editOrderInput.OrderId, validation.Required, is.UUIDv4),
		validation.Field(&editOrderInput.CustomerId, validation.Required, is.UUIDv4),
	)

	if err != nil {
		logger.Errorf("Invalid create order request | Error: %v | editOrderInput: %s", err.Error(), editOrderInput)
		return errors.New(errorcode.INVALID_INPUT_EDIT_ORDER)
	}

	productIdMap := make(map[string]struct{})

	for _, t := range editOrderInput.OrderItems {
		if _, ok := productIdMap[t.ProductId]; ok {
			logger.Error("Invalid edit order request duplicate product id")
			return errors.New(errorcode.INVALID_INPUT_EDIT_ORDER)
		}
		err = validation.ValidateStruct(&t,
			validation.Field(&t.ProductId, validation.Required, is.UUIDv4),
			validation.Field(&t.Quantity, validation.Required, validation.Min(0)),
		)

		if err != nil {
			logger.Errorf("Invalid create order request | Error: %v | editOrderInput: %s", err.Error(), editOrderInput)
			return errors.New(errorcode.INVALID_INPUT_EDIT_ORDER)
		}

		productIdMap[t.ProductId] = struct{}{}
	}

	return

}
