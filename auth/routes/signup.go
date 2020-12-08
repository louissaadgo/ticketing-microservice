package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"unicode"

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
	//Checks if the password is valid
	if invalid, errorMsg := isPasswordInvalid(credentials.Password); invalid {
		http.Error(w, errorMsg, http.StatusBadRequest)
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

//Checks if the password is Invalid
func isPasswordInvalid(s string) (valid bool, errorMsg string) {
	invalid := false
	number := false
	letter := false
	upperLetter := false
	specialLetter := false
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upperLetter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			specialLetter = true
		case unicode.IsLetter(c) || c == ' ':
			letter = true
		default:
			invalid = true
		}
	}
	if invalid == true {
		return true, "Invalid Password Unknown Reason"
	} else if number == false {
		return true, "Invalid Password - Password must contain a number"
	} else if letter == false {
		return true, "Invalid Password - Password must contain a letter"
	} else if upperLetter == false {
		return true, "Invalid Password - Password must contain an uppercase letter"
	} else if specialLetter == false {
		return true, "Invalid Password - Password must contain a special letter"
	} else if len(s) < 7 {
		return true, "Invalid Password - Password must be at least 7 characters"
	}
	return false, ""
}
