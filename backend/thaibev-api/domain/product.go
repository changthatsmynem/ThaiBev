package domain

import "time"

type ProductModel struct {
	ID        int       `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Barcode   string    `json:"barcode" db:"barcode"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ProductRequest struct {
	Code string `json:"code" binding:"required"`
}

type BarcodeResponse struct {
	Barcode string `json:"barcode"`
	Image   string `json:"image"`
}
