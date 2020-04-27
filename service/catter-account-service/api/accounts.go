package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielgospodinow/Catter/service/catter-account-service/models"
	"github.com/gorilla/mux"
)

// CreateAccount creates a new account.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	_ = json.NewDecoder(r.Body).Decode(&acc)

	err := models.CreateAccount(acc)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(acc)
	}
}

// GetAccount retrieves an account.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	acc, err := models.GetAccount(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(acc)
	}
}

// UpdateAccount updates an account.
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteAccount deletes an account.
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	acc, err := models.DeleteAccount(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(acc)
	}
}
