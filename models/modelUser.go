package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Age   string             `json:"age,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" validate:"required"`
}
