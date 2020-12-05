package routes

import "net/http"

//Signin signs in the user to our app
func Signin(w http.ResponseWriter, r *http.Request) {
	//These are equivalent to http.Error
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("500 - Something bad happened!"))
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}
