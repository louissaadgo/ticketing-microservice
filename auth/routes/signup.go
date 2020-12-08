package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/louissaadgo/ticketing-microservice/auth/user"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//Signup signs up the user to our app
func Signup(w http.ResponseWriter, r *http.Request) {
	credentials := user.Model{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Fatalln("Failed decoding json: ", err)
		return
	}
	//Checks if the email address is valid
	if isEmailInvalid(credentials.Email) {
		http.Error(w, "Inavlid Email Address", http.StatusBadRequest)
		return
	}
	if len(credentials.Password) < 5 {
		http.Error(w, "Password must be 5 characters or more", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Signed up successfully")
}

//Checks if the email is invalid
func isEmailInvalid(e string) bool {
	if len(e) < 5 || len(e) > 254 {
		return true
	}
	if emailRegex.MatchString(e) {
		return false
	}
	if !strings.Contains(e, "@") {
		return true
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return true
	}
	return false
}
