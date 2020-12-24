package routes

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

//Signout signs out the user from our app
func Signout(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
