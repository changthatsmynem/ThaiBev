package routes

import (
	"thaibev-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/ping", handlers.Ping)
		api.POST("/data", handlers.CreateData)
		api.DELETE("/data/:id", handlers.DeleteData)
	}
}