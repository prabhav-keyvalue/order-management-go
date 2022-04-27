package dto

import (
	"strconv"

	"github.com/gin-gonic/gin"
	defaults "github.com/prabhav-keyvalue/order-management-go/constant/default"
)

type PaginationParams struct {
	*QueryParamParser
	Offset int
	Limit  int
}

func NewPaginationParams() *PaginationParams {
	q := &QueryParamParser{}
	p := &PaginationParams{QueryParamParser: q}
	q.IQueryParamParser = p
	return p
}

func (p *PaginationParams) ParseQueryParams(c *gin.Context) (err error) {

	paginationParams := PaginationParams{}

	if l := c.Query("limit"); l != "" {
		limit, err := strconv.Atoi(l)
		if err != nil {
			return err
		}
		paginationParams.Limit = limit
	} else {
		paginationParams.Limit = defaults.PAGE_LIMIT_DEFAULT
	}

	if ofs := c.Query("offset"); ofs != "" {
		offset, err := strconv.Atoi(ofs)
		if err != nil {
			return err
		}
		paginationParams.Offset = offset
	} else {
		paginationParams.Offset = 0
	}

	c.Set("PaginationQueryParams", paginationParams)

	return
}
