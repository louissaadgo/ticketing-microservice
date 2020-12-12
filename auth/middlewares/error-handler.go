package middlewares

import (
	"fmt"
	"net/http"
)

//Error sends an error
func Error(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, err)
}
