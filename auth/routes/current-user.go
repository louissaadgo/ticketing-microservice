package routes

import (
	"fmt"
	"net/http"
)

//CurrentUser gets the current user signed to our app
func CurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi")
}
