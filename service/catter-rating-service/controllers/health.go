package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(context *gin.Context) {
	context.Status(http.StatusOK)
}
