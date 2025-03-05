package repositories

import (
	"context"

	"github.com/bybhub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

func CreateNewUser(db *mongo.Database, collectionName string, user *models.User) error {
	_, err := db.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(db *mongo.Database, collectionName string, filter bson.M) *models.UserResponse {
	var result *models.UserResponse
	info := db.Collection(collectionName).FindOne(ctx, filter).Decode(&result)
	if info != nil {
		return nil
	}
	return result
}
