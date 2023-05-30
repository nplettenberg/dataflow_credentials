package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credentials struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name,omitempty" `
	Secret string             `json:"secret" bson:"secret,omitempty" `
}

type CredentialsPreview struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty" `
}
