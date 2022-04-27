package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/prabhav-keyvalue/order-management-go/config"
	errorcode "github.com/prabhav-keyvalue/order-management-go/constant/error_code"
	"github.com/prabhav-keyvalue/order-management-go/entity"
	"github.com/prabhav-keyvalue/order-management-go/logger"

	repoMocks "github.com/prabhav-keyvalue/order-management-go/repository/mocks"
	"github.com/prabhav-keyvalue/order-management-go/service"
	serviceMocks "github.com/prabhav-keyvalue/order-management-go/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	orderRepository  repoMocks.OrderRepository
	orderItemService serviceMocks.OrderItemService
	productService   serviceMocks.ProductService
	service          service.OrderService
	order            entity.Order
}

func (suite *ServiceTestSuite) SetupSuite() {
	fmt.Println("hereeee")
	logger.InitLogger(config.Environment(config.GetEnv()))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.orderRepository = repoMocks.OrderRepository{}
	suite.orderItemService = serviceMocks.OrderItemService{}
	suite.productService = serviceMocks.ProductService{}
	suite.order = entity.Order{
		BaseEntity: entity.BaseEntity{
			Id: "id1",
		},
	}
	suite.service = &service.OrderServiceImpl{}
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.orderItemService.AssertExpectations(suite.T())
	suite.orderRepository.AssertExpectations(suite.T())
	suite.productService.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGetOrderById() {
	t := suite.T()
	t.Run("should successfully return order given order id", func(t *testing.T) {
		suite.SetupTest()
		defer suite.TearDownTest()
		orderId := "id"
		suite.orderRepository.On("GetOrderById", orderId).Return(suite.order, nil)
		res, err := suite.service.GetOrderById(orderId)

		assert.Equal(t, "fsdf", res.Id)
		require.NoError(t, err)
	})

	t.Run("should return errorcode ORDER_NOT_FOUND if order not found", func(t *testing.T) {
		suite.SetupTest()
		defer suite.TearDownTest()
		orderId := "id"
		suite.orderRepository.On("GetOrderById", orderId).Return(nil, errors.New("error"))
		_, err := suite.service.GetOrderById(orderId)

		require.NoError(t, err)
		assert.Equal(t, err.Error(), errorcode.ORDER_NOT_FOUND)
	})
}
