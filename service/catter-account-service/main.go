package main

import (
	"log"
	"net/http"

	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
	"github.com/danielgospodinow/Catter/service/catter-account-service/routers"
)

func main() {
	db.InitClient()
	defer db.CloseClient()

	router := routers.InitRouter()

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
