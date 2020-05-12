package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
	"github.com/danielgospodinow/Catter/service/catter-account-service/models"
	"github.com/danielgospodinow/Catter/service/catter-account-service/routers"
)

func main() {
	db.InitDynamoClient()
	defer db.CloseDynamoClient()

	models.InitAccountsTable()

	router := routers.InitRouter()

	http.Handle("/", router)
	fmt.Println("Server listening at port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
