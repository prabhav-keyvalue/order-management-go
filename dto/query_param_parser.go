package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/model"
	apiutil "github.com/prabhav-keyvalue/order-management-go/util/api"
)

type IQueryParamParser interface {
	ParseQueryParams(c *gin.Context) error
	Parse(c *gin.Context)
}

type QueryParamParser struct {
	IQueryParamParser
}

func (qp *QueryParamParser) Parse(c *gin.Context) {
	err := qp.ParseQueryParams(c)

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
