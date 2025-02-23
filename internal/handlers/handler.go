package handler

import (
	"github.com/bybhub/backend/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger *config.Logger
	db     *mongo.Database
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMongoDB()
}
