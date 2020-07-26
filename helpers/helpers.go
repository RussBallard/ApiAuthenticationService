package helpers

import (
	"net/http"
)

func SetResponseHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	return w
}
