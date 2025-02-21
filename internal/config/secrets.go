package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitializeSecrets() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func GetEnv() string {
	return os.Getenv("ENV")
}

func GetPort() string {
	return os.Getenv("PORT")
}
