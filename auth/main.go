package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/louissaadgo/ticketing-microservice/auth/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Port 3000
const address string = ":3000"

//Client all
var Client *mongo.Client
var err error

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://auth-mongo-srv:27017/auth"))
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(ctx)
	r := mux.NewRouter()
	r.HandleFunc("/api/users/signup", signup).Methods("POST")
	r.HandleFunc("/api/users/signin", signin).Methods("POST")
	r.HandleFunc("/api/users/signout", signout).Methods("POST")
	r.HandleFunc("/api/users/currentuser", currentUser).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}

func signup(w http.ResponseWriter, r *http.Request) {
	routes.Signup(w, r, Client)
}
func signin(w http.ResponseWriter, r *http.Request) {
	routes.Signin(w, r, Client)
}
func signout(w http.ResponseWriter, r *http.Request) {
	routes.Signout(w, r, Client)
}
func currentUser(w http.ResponseWriter, r *http.Request) {
	routes.CurrentUser(w, r, Client)
}
