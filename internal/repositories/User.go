package repositories

import (
	"context"
	"errors"

	"github.com/bybhub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

// criando um novo usuário
func CreateNewUser(db *mongo.Database, collectionName string, user *models.User) error {
	_, err := db.Collection(collectionName).InsertOne(ctx, user)
	return err
}

// buscar usuario pelo ID
func FindUserByID(db *mongo.Database, collectionName string, id string) (*models.UserResponse, error) {
	var result models.UserResponse
	err := db.Collection(collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}
	return &result, nil
}

// atualizar usuario pelo id
func UpdateUserByID(db *mongo.Database, collectionName string, id string, update bson.M) error {
	_, err := db.Collection(collectionName).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

// deletar usuario pelo id
func DeleteUserByID(db *mongo.Database, collectionName string, id string) error {
	_, err := db.Collection(collectionName).DeleteOne(ctx, bson.M{"_id": id})
	return err
}
