package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	errortype "github.com/louissaadgo/ticketing-microservice/auth/errorType"
)

//ErrorHandler sends an error
func ErrorHandler(w http.ResponseWriter, errors errortype.Universal, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	bs, err := json.Marshal(errors)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	fmt.Fprint(w, string(bs))
}
