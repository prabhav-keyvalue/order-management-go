package server

import (
	"github.com/prabhav-keyvalue/order-management-go/config"
	"github.com/prabhav-keyvalue/order-management-go/db"
	"github.com/prabhav-keyvalue/order-management-go/logger"
	"github.com/prabhav-keyvalue/order-management-go/server/router"
)

func Start() (err error) {
	err = config.Load(".env")

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

	err = r.Router.Run(":" + config.GetAppPort())

	if err != nil {
		return
	}

	return
}
