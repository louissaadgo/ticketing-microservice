package routes

import "net/http"

//CurrentUser gets the current user signed to our app
func CurrentUser(w http.ResponseWriter, r *http.Request) {
	//These are equivalent to http.Error
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("500 - Something bad happened!"))
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
