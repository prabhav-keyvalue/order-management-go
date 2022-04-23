package apiutil

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/model"
)

func SendResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, &model.Response{
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, errorCode string, message string) {
	c.JSON(statusCode, &model.Response{
		Error: &model.ApiError{
			Code:    errorCode,
			Message: message,
		},
	})
}
