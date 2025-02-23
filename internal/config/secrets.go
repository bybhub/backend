package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ENV_FILE = ".env"
)

type Secrets struct {
	ENV                 string `env:"ENV"`
	REDIS_URL           string `env:"REDIS_URL"`
	MONGO_DB_USER       string `env:"MONGO_DB_USER"`
	MONGO_DB_PASS       string `env:"MONGO_DB_PASS"`
	MONGO_DB_HOST       string `env:"MONGO_DB_HOST"`
	MONGO_DB_HOST_LOCAL string `env:"MONGO_DB_HOST_LOCAL"`
	SECRET_KEY_JWT      string `env:"SECRET_KEY_JWT"`
}

func InitializeSecrets() *Secrets {
	logger = GetLogger("secrets")
	secrets := parseEnv()
	return secrets
}

func loadEnvFile() error {
	return godotenv.Load(ENV_FILE)
}

func parseEnv() *Secrets {
	if err := loadEnvFile(); err != nil {
		logger.Errorf("erro to load dotenv file, error: %v", err)
		return &Secrets{}
	}
	sct := &Secrets{}
	if err := env.Parse(sct); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return sct
}
