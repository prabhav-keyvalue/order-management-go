// @title           Order Management API
// @version         1.0
// @description     Order management rest api server made with Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath /api/v1
package main

import (
	"fmt"
	"os"

	"github.com/prabhav-keyvalue/order-management-go/server"
)

func main() {
	err := server.Start()

	if err != nil {
		fmt.Println("Server Unable to start, Error: ", err.Error())
		os.Exit(1)
	}

}
