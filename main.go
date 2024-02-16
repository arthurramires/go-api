package main

import (
	"github.com/arthurramires/go-api/config"
	"github.com/arthurramires/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
