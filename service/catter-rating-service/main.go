package main

import (
	"github.com/danielgospodinow/Catter/service/catter-rating-service/routers"
)

func main() {
	router := routers.InitRouter()

	_ = router.Run()
}