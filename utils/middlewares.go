package utils

import (
	"fmt"
	"net/http"
)

func DeleteMethodHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if tokenApi := r.Header.Get("X-TokenApi"); tokenApi != "JKUJJ718VRMQDNIMZZL3" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "%s"}`, "authentication token not passed")
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
