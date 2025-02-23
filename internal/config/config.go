package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db      *mongo.Database
	logger  *Logger
	secrets *Secrets
)

func Init() error {
	secrets = InitializeSecrets()

	var err error
	db, err = InitalizeDatabase("byb-db")
	if err != nil {
		return fmt.Errorf("erro in mongodb init: %w", err)
	}
	return nil
}

func GetMongoDB() *mongo.Database {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetSecrets() *Secrets {
	return secrets
}
