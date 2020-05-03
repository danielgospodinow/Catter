package api

import (
	"net/http"
)

// CheckHealth returns a plain HTTP 200 response, just to signal that
// the service is working properly.
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
