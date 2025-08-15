package handlers

import (
	"net/http"
	"thaibev-api/models"
	"thaibev-api/services"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		Message: "pong",
	})
}

func CreateData(c *gin.Context) {
	var req models.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid JSON",
		})
		return
	}

	// Generate barcode
	barcodeService := services.NewBarcodeService()
	barcodeImage, err := barcodeService.GenerateBarcode(req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to generate barcode",
		})
		return
	}

	product := models.ProductModel{
		Code:    req.Code,
		Barcode: barcodeImage,
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Product created with barcode",
		Data: models.BarcodeResponse{
			Barcode: product.Barcode,
			Image:   barcodeImage,
		},
	})
}

func DeleteData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Data deleted",
		Data:    gin.H{"id": id},
	})
}
