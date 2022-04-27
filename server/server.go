package server

import (
	"github.com/prabhav-keyvalue/order-management-go/config"
	"github.com/prabhav-keyvalue/order-management-go/db"
	docs "github.com/prabhav-keyvalue/order-management-go/docs"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/server/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() (err error) {
	err = config.Load(".env")

	docs.SwaggerInfo.Title = "Order Management documentation"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api/v1"

	if err != nil {
		return
	}

	_, err = logger.InitLogger(config.Environment(config.GetEnv()))

	if err != nil {
		return
	}

	err = db.InitDb()

	if err != nil {
		return
	}

	r := router.InitRouter()

	if config.GetEnv() != string(config.Production) {
		r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	err = r.Router.Run(":" + config.GetAppPort())

	if err != nil {
		return
	}

	return
}
