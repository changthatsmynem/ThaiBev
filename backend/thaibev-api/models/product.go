package models

import "thaibev-api/domain"

type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    []*domain.ProductModel `json:"data,omitempty"`
}
