package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		apiutil.SendErrorResponse(c, http.StatusBadRequest, err.Error(), "Invalid input")
		c.Abort()
		return
	}

	c.Next()
}
