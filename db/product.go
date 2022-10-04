package db

import (
	"database/sql"
	"github.com/madz0k/demo-crm-backend/models"
)

func (db Database) GetAllProducts() (*models.ProductList, error) {
	list := &models.ProductList{}

	rows, err := db.Conn.Query("SELECT * FROM products ORDER BY ID DESC;")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Products = append(list.Products, product)
	}
	return list, nil
}

func (db Database) AddProduct(product *models.Product) error {
	var id int
	var createdAt string
	query := `INSERT INTO products (name) VALUES ($1) RETURNING id, createdAt;`
	err := db.Conn.QueryRow(query, product.Name).Scan(&id, &createdAt)
	if err != nil {
		return nil
	}

	product.ID = id
	product.CreatedAt = createdAt

	return nil
}

func (db Database) GetProductById(productId int) (models.Product, error) {
	product := models.Product{}

	query := `SELECT * FROM products WHERE id = $1;`
	row := db.Conn.QueryRow(query, productId)
	switch err := row.Scan(&product.ID, &product.Name, &product.CreatedAt); err {
	case sql.ErrNoRows:
		return product, ErrNoMatch
	default:
		return product, err
	}
}

func (db Database) DeleteProduct(productId int) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := db.Conn.Exec(query, productId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateProduct(productId int, productData models.Product) (models.Product, error) {
	product := models.Product{}
	query := `UPDATE products SET name=$1 WHERE id=$2 RETURNING id, name, created_at;`
	err := db.Conn.QueryRow(query, productData.Name, productId).Scan(&product.ID, &product.Name, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, ErrNoMatch
		}
		return product, err
	}
	return product, nil
}
