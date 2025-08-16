package repository

import "thaibev-api/domain"

type ProductRepository interface {
	GetAll() ([]*domain.ProductModel, error)
	Create(product *domain.ProductModel) error
	Delete(id int) error
}
