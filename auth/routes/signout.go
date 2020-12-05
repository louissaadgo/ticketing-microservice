package routes

import "net/http"

//Signout signs out the user from our app
func Signout(w http.ResponseWriter, r *http.Request) {
	//These are equivalent to http.Error
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("500 - Something bad happened!"))
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
