package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SortOptions struct {
	*QueryParamParser
	SortKey   string
	SortOrder string
}

func NewSortOptions() *SortOptions {
	q := &QueryParamParser{}
	s := &SortOptions{QueryParamParser: q}
	q.IQueryParamParser = s

	return s
}

func (s *SortOptions) ParseQueryParams(c *gin.Context) (err error) {
	sortOptions := SortOptions{}

	if qp := c.Query("sort_key"); qp != "" {
		err = validation.Validate(qp, validation.In("total_quantity", "total_price"))
		if err != nil {
			return
		}
		sortOptions.SortKey = qp
	} else {
		sortOptions.SortKey = "created_at"
	}

	if qp := c.Query("sort_order"); qp != "" {
		err = validation.Validate(qp, validation.In("asc", "desc"))
		if err != nil {
			return
		}
		sortOptions.SortOrder = qp
	} else {
		sortOptions.SortOrder = "desc"
	}

	fmt.Println("sdfdsfdsfsdfsd", sortOptions)

	c.Set("SortOptionsQueryParams", sortOptions)
	return
}
