package routes

import "net/http"

//CurrentUser gets the current user signed to our app
func CurrentUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
