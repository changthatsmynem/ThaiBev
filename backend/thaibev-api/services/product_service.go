package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"thaibev-api/domain"
	"thaibev-api/models"
	"thaibev-api/repository"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
)

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(body *domain.ProductRequest) (*models.Response, error) {
	barcode, err := s.GenerateBarcode(body.Code)
	if err != nil {
		return nil, err
	}

	product := &domain.ProductModel{
		Code:      body.Code,
		Barcode:   barcode,
		CreatedAt: time.Now(),
	}

	err = s.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	response := &models.Response{
		Status:  "success",
		Message: "Product created with barcode",
	}

	return response, nil
}

func (s *productService) GetAllProducts() (*models.Response, error) {
	res, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    res,
	}, nil
}

func (s *productService) DeleteProduct(id int) error {
	if err := s.productRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *productService) GenerateBarcode(productCode string) (string, error) {
	barcodeData, err := code39.Encode(productCode, false, true)
	if err != nil {
		fmt.Printf("Error encoding barcode: %v\n", err)
		return "", fmt.Errorf("failed to encode barcode: %w", err)
	}

	scaledBarcode, err := barcode.Scale(barcodeData, 300, 30)
	if err != nil {
		fmt.Printf("Error scaling barcode: %v\n", err)
		return "", fmt.Errorf("failed to scale barcode: %w", err)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, scaledBarcode)
	if err != nil {
		fmt.Printf("Error encoding PNG: %v\n", err)
		return "", fmt.Errorf("failed to encode PNG: %w", err)
	}

	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Image, nil
}
