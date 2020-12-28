package user

import "go.mongodb.org/mongo-driver/bson/primitive"

//Model is the user's authentication model
type Model struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
}
