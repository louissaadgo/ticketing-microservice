package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//structure is the error structure
type structure struct {
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`
}

//ErrorHandler sends an error
func ErrorHandler(w http.ResponseWriter, errorType string, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, toJSON(errorType, message))
}

//toJSON transforms the struct into json
func toJSON(errorType string, message string) string {
	errorStructure := structure{
		ErrorType: errorType,
		Message:   message,
	}
	bs, err := json.Marshal(errorStructure)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	return string(bs)
}
