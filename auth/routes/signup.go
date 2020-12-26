package routes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	errortype "github.com/louissaadgo/ticketing-microservice/auth/errorType"
	"github.com/louissaadgo/ticketing-microservice/auth/middlewares"
	"github.com/louissaadgo/ticketing-microservice/auth/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//Signup signs up the user to our app
func Signup(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	credentials := user.Model{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Fatalln("Failed decoding json: ", err)
		return
	}
	//Checks if the email address is invalid
	if invalid, allErrors := isEmailInvalid(credentials.Email); invalid {
		newError := errortype.RequestValidationError{
			Errors: allErrors,
		}
		newErrors := []errortype.ErrorModel{}
		for _, a := range newError.Errors {
			newErrors = append(newErrors, errortype.ErrorModel{
				Field:   "Email",
				Message: a,
			})
		}
		middlewares.ErrorHandler(w, newErrors, http.StatusBadRequest)
		return
	}
	//Checks if the password is invalid
	if invalid, allErrors := isPasswordInvalid(credentials.Password); invalid {
		newError := errortype.RequestValidationError{
			Errors: allErrors,
		}
		newErrors := []errortype.ErrorModel{}
		for _, a := range newError.Errors {
			newErrors = append(newErrors, errortype.ErrorModel{
				Field:   "Password",
				Message: a,
			})
		}
		middlewares.ErrorHandler(w, newErrors, http.StatusBadRequest)
		return
	}
	collection := client.Database("auth").Collection("users")
	//Checks if the email is already in use
	filter := bson.M{"email": credentials.Email}
	var check bson.M
	if err = collection.FindOne(context.TODO(), filter).Decode(&check); err == nil {
		newError := errortype.ErrorModel{Field: "Email", Message: "Email Already in use"}
		errorEmail := []errortype.ErrorModel{newError}
		middlewares.ErrorHandler(w, errorEmail, http.StatusBadRequest)
		return
	}
	hashedPassword := sha256.Sum256([]byte(credentials.Password + credentials.ID.Hex()))
	credentials.Password = hex.EncodeToString(hashedPassword[:])
	insertResult, err := collection.InsertOne(context.TODO(), credentials)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "SUCCESS SIGNUP - ID: ", insertResult.InsertedID)
}

//Checks if the email is invalid
func isEmailInvalid(e string) (bool, []string) {
	errors := []string{}
	invalid := false
	if len(e) < 5 || len(e) > 254 {
		errors = append(errors, "Email must be more than 4 characters and less than 254")
		invalid = true
	}
	if !emailRegex.MatchString(e) {
		invalid = true
	}
	if !strings.Contains(e, "@") {
		invalid = true
		errors = append(errors, "Email must contain domain name")
		return invalid, errors
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		errors = append(errors, "Invalid email domain DNS MX")
		invalid = true
	}
	return invalid, errors
}

//Checks if the password is Invalid
func isPasswordInvalid(s string) (bool, []string) {
	errors := []string{}
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
	if number == false {
		errors = append(errors, "Password must contain a number")
		invalid = true
	}
	if letter == false {
		errors = append(errors, "Password must contain a letter")
		invalid = true
	}
	if upperLetter == false {
		errors = append(errors, "Password must contain an uppercase letter")
		invalid = true
	}
	if specialLetter == false {
		errors = append(errors, "Password must contain a special letter")
		invalid = true
	}
	if len(s) < 7 {
		errors = append(errors, "Password must be at least 7 characters")
		invalid = true
	}
	return invalid, errors
}
