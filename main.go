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
