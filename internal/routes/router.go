package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	trustedProxies := []string{"127.0.0.1", "::1"}

	router := gin.Default()
	initializeRoutes(router)

	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Erro ao configurar proxies confiáveis: %v", err)
	}

	router.Run(":9999")
}
