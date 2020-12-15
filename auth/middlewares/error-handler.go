package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	errortype "github.com/louissaadgo/ticketing-microservice/auth/errorType"
)

//ReqValErrorHandler sends an error
func ReqValErrorHandler(w http.ResponseWriter, reqVal errortype.RequestValidationError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	bs, err := json.Marshal(reqVal)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	fmt.Fprint(w, string(bs))
}

//DBConnErrorHandler sends an error
func DBConnErrorHandler(w http.ResponseWriter, DBConn errortype.DatabaseConnectionError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	bs, err := json.Marshal(DBConn)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	fmt.Fprint(w, string(bs))
}
