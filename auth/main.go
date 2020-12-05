package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louissaadgo/ticketing-microservice/auth/routes"
)

//Port 3000
const address string = ":3000"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/signup", routes.Signup).Methods("POST")
	r.HandleFunc("/api/users/signin", routes.Signin).Methods("POST")
	r.HandleFunc("/api/users/signout", routes.Signout).Methods("POST")
	r.HandleFunc("/api/users/currentuser", routes.CurrentUser).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}
