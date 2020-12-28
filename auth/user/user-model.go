package user

import "go.mongodb.org/mongo-driver/bson/primitive"

//Model is the user's authentication model
type Model struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
