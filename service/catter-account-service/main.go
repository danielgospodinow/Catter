package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
	"github.com/danielgospodinow/Catter/service/catter-account-service/models"
	"github.com/danielgospodinow/Catter/service/catter-account-service/routers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDynamoClient()
	defer db.CloseDynamoClient()

	models.InitAccountsTable()

	router := routers.InitRouter()

	http.Handle("/", router)
	fmt.Println("Server listening at port 8080 ...")

	log.Fatal(http.ListenAndServe(":8080", enableCORS(router)))
}

func enableCORS(router *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(headersOk, originsOk, methodsOk)(router)
}
