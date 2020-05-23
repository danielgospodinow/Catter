package routers

import (
	"github.com/danielgospodinow/Catter/service/catter-rating-service/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router.
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", api.HealthCheck)

	return router
}
