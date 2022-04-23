package router

import "github.com/gin-gonic/gin"

type Router struct {
	Router *gin.Engine
}

func InitRouter() Router {
	r := Router{
		Router: gin.New(),
	}

	r.Router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)

	root := r.Router.Group("/api")
	r.initOrderRoutes(root)

	return r
}
