package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/dto"
)

func ParseQueryParams(parser *dto.QueryParamParser) gin.HandlerFunc {
	return parser.Parse
}
