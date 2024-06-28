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

	if err != nil {
		logger.Errof("config initializer erro: %v", err)
		return
	}

	router.Initializer()
}
