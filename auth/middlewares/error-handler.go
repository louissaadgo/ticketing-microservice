package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	errortype "github.com/louissaadgo/ticketing-microservice/auth/errorType"
)

//ReqValErrorHandler sends an error
func ReqValErrorHandler(w http.ResponseWriter, reqVal errortype.RequestValidationError, param string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	newError := errortype.Universal{}
	for _, a := range reqVal.Errors {
		newError.Errors = append(newError.Errors, errortype.ErrorModel{
			Field:   param,
			Message: a,
		})
	}
	bs, err := json.Marshal(newError)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	fmt.Fprint(w, string(bs))
}

//DBConnErrorHandler sends an error
func DBConnErrorHandler(w http.ResponseWriter, DBConn errortype.DatabaseConnectionError, param string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	newError := errortype.Universal{}
	newError.Errors = append(newError.Errors, errortype.ErrorModel{
		Field:   param,
		Message: DBConn.Error,
	})
	bs, err := json.Marshal(newError)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	fmt.Fprint(w, string(bs))
}
