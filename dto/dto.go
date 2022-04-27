package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/model"
	apiutil "github.com/prabhav-keyvalue/order-management-go/util/api"
)

type ISDto interface {
	Validate(c *gin.Context)
	validateDto(c *gin.Context) error
}

type Dto struct {
	ISDto
}

func (d *Dto) Validate(c *gin.Context) {
	err := d.validateDto(c)

	if err != nil {
		apiutil.SendResponse(c, http.StatusBadRequest, &model.Response{
			Error: &model.ApiError{
				Code: err.Error(),
			},
		})
		c.Abort()
		return
	}

	c.Next()
}
