package main

import (
	"github.com/wendellast/Maybe/config"
	"github.com/wendellast/Maybe/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Config Initializer
	err := config.Init()

	router.Initializer()

	if err != nil {
		logger.Errof("config initializer erro: %v", err)
		return
	}
}
