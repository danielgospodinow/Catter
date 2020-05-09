package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is used to determine the health status of the service.
func HealthCheck(context *gin.Context) {
	context.Status(http.StatusOK)
}
