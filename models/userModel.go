package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  *string            `json:"name" bson:"_name"`
	Email *string            `json:"email" bson:"_email"`
}
