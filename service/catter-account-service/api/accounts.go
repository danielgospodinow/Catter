package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/danielgospodinow/Catter/service/catter-account-service/models"
	"github.com/gorilla/mux"
)

// CreateAccount creates a new account.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	_ = json.NewDecoder(r.Body).Decode(&acc)

	racc, err := models.CreateAccount(acc)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(racc)
	}
}

// GetAccount retrieves an account.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]

	acc, err := models.GetAccount(email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	if reflect.DeepEqual(models.Account{}, acc) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(acc)
}

// UpdateAccount updates an account.
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteAccount deletes an account.
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]

	acc, err := models.DeleteAccount(email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(acc)
	}
}
