package routers

import (
	"net/http"

	"github.com/danielgospodinow/Catter/service/catter-account-service/api"
	"github.com/gorilla/mux"
)

// InitRouter initializes router and returns an instance to it.
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", api.CheckHealth).Methods(http.MethodGet)

	router.HandleFunc("/account", api.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/account/{id}", api.GetAccount).Methods(http.MethodGet)
	router.HandleFunc("/account/{id}", api.UpdateAccount).Methods(http.MethodPut)
	router.HandleFunc("/account/{id}", api.DeleteAccount).Methods(http.MethodDelete)

	return router
}
