package routers

import (
	"github.com/danielgospodinow/Catter/service/catter-rating-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", controllers.HealthCheck)

	return router
}
