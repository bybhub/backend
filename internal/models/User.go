package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

type UserResponse struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}
