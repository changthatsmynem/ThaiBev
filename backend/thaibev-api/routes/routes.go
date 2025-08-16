package routes

import (
	"thaibev-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, productHandler *handlers.ProductHandler) {
	api := r.Group("/api/products")
	{
		api.POST("", productHandler.CreateProduct)
		api.GET("", productHandler.GetAllProducts)
		api.DELETE(":id", productHandler.DeleteProduct)
	}
}
