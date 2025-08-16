package services

import (
	"thaibev-api/domain"
	"thaibev-api/models"
)

type ProductService interface {
	GetAllProducts() (*models.Response, error)
	CreateProduct(body *domain.ProductRequest) (*models.Response, error)
	DeleteProduct(id int) error
	GenerateBarcode(productCode string) (string, error)
}
