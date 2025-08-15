package models

import "time"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

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
	Image   string `json:"image"` // base64 encoded barcode image
}
