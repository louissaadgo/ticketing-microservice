package routes

import (
	"encoding/json"
	"net/http"

	"github.com/louissaadgo/ticketing-microservice/auth/user"
)

//Signin signs in the user to our app
func Signin(w http.ResponseWriter, r *http.Request) {
	credentials := user.Model{}
	json.NewDecoder(r.Body).Decode(&credentials)
	//Check credentials
}
