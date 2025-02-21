package main

import (
	"github.com/bybhub/backend/internal/config"
	"github.com/bybhub/backend/internal/routes"
)

func main() {
	config.InitializeSecrets()
	routes.InitializeRouter()
}
