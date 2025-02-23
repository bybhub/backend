package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name  string
	Email string
}

type UserResponse struct {
	ID    primitive.ObjectID `json:"_id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}
