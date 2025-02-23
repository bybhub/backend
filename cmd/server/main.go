package main

import (
	"github.com/bybhub/backend/internal/config"
	"github.com/bybhub/backend/internal/routes"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config init error: %s", err)
		return
	}
	routes.InitializeRouter()
}
