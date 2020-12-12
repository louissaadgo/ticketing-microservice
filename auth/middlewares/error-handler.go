package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//structure is the error structure
type structure struct {
	Message string `json:"message"`
}

//ErrorHandler sends an error
func ErrorHandler(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, toJSON(message))
}

//toJSON transforms the struct into json
func toJSON(message string) string {
	errorStructure := structure{
		Message: message,
	}
	bs, err := json.Marshal(errorStructure)
	if err != nil {
		log.Fatalln("Failed to marshal JSON")
	}
	return string(bs)
}
