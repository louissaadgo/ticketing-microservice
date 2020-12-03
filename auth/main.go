package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 3000
const address string = ":3000"

//Sends the current user
func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	//These are equivalent to http.Error
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("500 - Something bad happened!"))
	http.Error(w, "Invalid email adress", http.StatusForbidden)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/currentuser", getCurrentUser).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}
