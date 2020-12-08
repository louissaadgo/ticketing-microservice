package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/louissaadgo/ticketing-microservice/auth/user"
)

//Signup signs up the user to our app
func Signup(w http.ResponseWriter, r *http.Request) {
	credentials := user.Model{}
	json.NewDecoder(r.Body).Decode(&credentials)
	if len(credentials.Email) < 10 {
		http.Error(w, "Email must be longer than 10 characters", http.StatusBadRequest)
		return
	}
	if len(credentials.Password) < 5 {
		http.Error(w, "Password must be 5 characters or more", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Signed up successfully")
}
