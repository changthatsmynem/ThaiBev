package handlers

import (
	"net/http"
	"strconv"
	"thaibev-api/domain"
	"thaibev-api/models"
	"thaibev-api/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req *domain.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	res, err := h.productService.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	res, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	err = h.productService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Product deleted successfully",
	})
}
