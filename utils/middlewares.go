package utils

import (
	"mime"
	"net/http"
)

func DeleteMethodHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "" {
			http.Error(w, "Bad syntax", http.StatusBadRequest)
			return

		}

		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil || mt != "application/json" {
			http.Error(w, "Bad syntax", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
