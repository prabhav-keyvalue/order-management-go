package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/dto"
)

func ValidateDto(dto dto.Dto) gin.HandlerFunc {
	return dto.Validate
}
