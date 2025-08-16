package repository

import (
	"database/sql"
	"thaibev-api/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return productRepository{db: db}
}

func (r productRepository) Create(product *domain.ProductModel) error {
	query := `INSERT INTO products (code, barcode, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, product.Code, product.Barcode, product.CreatedAt)
	return err
}

func (r productRepository) GetAll() ([]*domain.ProductModel, error) {
	query := `SELECT id, code, barcode, created_at FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.ProductModel
	for rows.Next() {
		product := &domain.ProductModel{}
		err := rows.Scan(&product.ID, &product.Code, &product.Barcode, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r productRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
