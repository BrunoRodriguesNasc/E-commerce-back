package repositories

import (
	"database/sql"
	"e-commerce/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, rating, size, reviews, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	return r.db.QueryRow(query, product.Name, product.Description, product.Price, product.Rating, product.Size, product.Reviews, product.Image	).Scan(&product.ID)
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	products := []models.Product{}
	rows, err := r.db.Query(`SELECT id, name, description, price, rating, size, reviews, image, created_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Rating, &p.Size, &p.Reviews, &p.Image, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
