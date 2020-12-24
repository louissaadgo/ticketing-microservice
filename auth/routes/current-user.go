package routes

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

//CurrentUser gets the current user signed to our app
func CurrentUser(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	fmt.Fprintln(w, "Hi")
}
