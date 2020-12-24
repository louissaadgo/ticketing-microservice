package routes

import (
	"encoding/json"
	"net/http"

	"github.com/louissaadgo/ticketing-microservice/auth/user"
	"go.mongodb.org/mongo-driver/mongo"
)

//Signin signs in the user to our app
func Signin(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	credentials := user.Model{}
	json.NewDecoder(r.Body).Decode(&credentials)
	//Check credentials
}
