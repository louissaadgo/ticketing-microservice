package routes

import "net/http"

//Signout signs out the user from our app
func Signout(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
