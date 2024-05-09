package models

import (
	"context"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Insert(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int

	stmt := `INSERT INTO products (name, description, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5) returning id`

	err := db.QueryRowContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil

}

func (p *Product) GetAll() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, description, price, created_at, updated_at FROM products ORDER BY name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *Product) GetOneById(id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil

}

func (p *Product) Update(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE products SET
		name = $1,
		description = $2,
		price = $3,
		updated_at = $4
		where id = $5`

	_, err := db.ExecContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		product.Id,
	)

	if err != nil {
		return 0, err
	}

	return 0, nil

}

func (p *Product) DeleteByID(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE FROM products WHERE id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
