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
	r.HandleFunc("/api/users/signup", signUp).Methods("POST")
	r.HandleFunc("/api/users/signin", routes.Signin).Methods("POST")
	r.HandleFunc("/api/users/signout", routes.Signout).Methods("POST")
	r.HandleFunc("/api/users/currentuser", routes.CurrentUser).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	routes.Signup(w, r, Client)
}
