package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 3000
const address string = ":3000"

//Sends the current user
func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/currentuser", getCurrentUser).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}
