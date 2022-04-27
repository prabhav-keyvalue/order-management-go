package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhav-keyvalue/order-management-go/controller"
	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/dto"
	"github.com/prabhav-keyvalue/order-management-go/middleware"
)

func (r Router) initOrderRoutes(rg *gin.RouterGroup) {
	db := db.GetDB()
	orderController := controller.NewOrderController(db)
	v1 := rg.Group("/v1")
	orderV1 := v1.Group("/orders")
	{
		orderV1.PUT("", middleware.ValidateDto(dto.NewEditOrderInputDto().Dto), orderController.EditOrder)
		orderV1.POST("", middleware.ValidateDto(dto.NewCreateOrderInputDto().Dto), orderController.CreateOrder)
		orderV1.GET("/:id", orderController.GetOrderById)
		orderV1.DELETE("/:id", orderController.DeleteOrder)
		orderV1.GET("",
			middleware.ParseQueryParams(dto.NewPaginationParams().QueryParamParser),
			middleware.ParseQueryParams(dto.NewSortOptions().QueryParamParser),
			middleware.ParseQueryParams(dto.NewOrderFilterParams().QueryParamParser),
			orderController.GetOrders)
	}
}
