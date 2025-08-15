package handlers

import (
	"net/http"
	"thaibev-api/models"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		Status:  "ok",
		Message: "ThaiBev API is running",
	})
}