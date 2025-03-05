package repositories

import (
	"context"
	"errors"

	"github.com/bybhub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

// criando um novo usuário
func CreateNewUser(db *mongo.Database, collectionName string, user *models.User) (primitive.ObjectID, error) {
	result, err := db.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func FindAllUsers(db *mongo.Database, collectionName string) ([]models.UserResponse, error) {
	var results []models.UserResponse
	cursor, err := db.Collection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result models.UserResponse
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// buscar usuario pelo ID
func FindUserByID(db *mongo.Database, collectionName string, id string) (*models.UserResponse, error) {
	var result models.UserResponse

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	return &result, nil
}

func UpdateUserByID(db *mongo.Database, collectionName string, id string, update bson.M) (*models.UserResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	result, err := db.Collection(collectionName).UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("usuário não encontrado")
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("nenhuma alteração realizada, dados já estavam atualizados")
	}

	var updatedUser models.UserResponse
	err = db.Collection(collectionName).FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&updatedUser)
	if err != nil {
		return nil, errors.New("erro ao buscar usuário atualizado")
	}

	return &updatedUser, nil
}

func DeleteUserByID(db *mongo.Database, collectionName string, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	result, err := db.Collection(collectionName).DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("usuário não encontrado")
	}

	return nil
}
