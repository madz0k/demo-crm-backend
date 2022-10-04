package db

import "github.com/madz0k/demo-crm-backend/models"

func (db Database) GetAllProducts() (*models.ProductList, error) {
	list := &models.ProductList{}

	rows, err := db.Conn.Query("SELECT * FROM products ORDER BY ID DESC")
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
