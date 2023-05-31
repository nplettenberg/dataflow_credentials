package main

type Credentials struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty" `
	Secret string `json:"secret" bson:"secret,omitempty" `
}

type CredentialsPreview struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name,omitempty" `
}
