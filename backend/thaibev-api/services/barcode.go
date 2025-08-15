package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

type BarcodeService struct{}

func NewBarcodeService() *BarcodeService {
	return &BarcodeService{}
}

func (s *BarcodeService) GenerateBarcode(productCode string) (string, error) {
	barcodeData, err := code128.Encode(productCode)
	if err != nil {
		return "", fmt.Errorf("failed to encode barcode: %w", err)
	}

	scaledBarcode, err := barcode.Scale(barcodeData, 200, 100)
	if err != nil {
		return "", fmt.Errorf("failed to scale barcode: %w", err)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, scaledBarcode)
	if err != nil {
		return "", fmt.Errorf("failed to encode PNG: %w", err)
	}

	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Image, nil
}
