package dto

import (
	"github.com/gin-gonic/gin"
)

type OrderFilterParams struct {
	*QueryParamParser
	MaxQuantity string
	MinQuantity string
}

func NewOrderFilterParams() *OrderFilterParams {
	q := &QueryParamParser{}
	o := &OrderFilterParams{QueryParamParser: q}
	q.IQueryParamParser = o
	return o
}

func (o *OrderFilterParams) ParseQueryParams(c *gin.Context) (err error) {
	orderFilterParams := OrderFilterParams{}

	if qp := c.Query("max_quantity"); qp != "" {

		orderFilterParams.MaxQuantity = qp
	}

	if qp := c.Query("min_quantity"); qp != "" {
		orderFilterParams.MinQuantity = qp
	}

	c.Set("orderFilterParams", orderFilterParams)

	return
}
