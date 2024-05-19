package main

import (
	"fmt"

	"github.com/wendellast/Maybe/config"
	"github.com/wendellast/Maybe/router"
)

func main() {
	//Config Initializer
	err := config.Init()

	router.Initializer()

	if err != nil {
		fmt.Println(err)
		return
	}
}
