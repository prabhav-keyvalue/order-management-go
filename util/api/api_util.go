package apiutil

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/model"
)

func SendResponse(c *gin.Context, statusCode int, data *model.Response) {
	c.JSON(statusCode, data)
}
