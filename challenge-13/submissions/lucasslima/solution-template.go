package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Product represents a product in the inventory system
type Product struct {
	ID       int64
	Name     string
	Price    float64
	Quantity int
	Category string
}

// ProductStore manages product operations
type ProductStore struct {
	db *sql.DB
}

// NewProductStore creates a new ProductStore with the given database connection
func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

// InitDB sets up a new SQLite database and creates the products table
func InitDB(dbPath string) (*sql.DB, error) {
	// TODO: Open a SQLite database connection
	// TODO: Create the products table if it doesn't exist
	// The table should have columns: id, name, price, quantity, category
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, name varchar(50), price float, quantity int, category varchar)")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateProduct adds a new product to the database
func (ps *ProductStore) CreateProduct(product *Product) error {
	// TODO: Insert the product into the database
	// TODO: Update the product.ID with the database-generated ID
	result, err := ps.db.Exec("INSERT INTO products (name,price,quantity,category) VALUES (?,?,?,?)", product.Name, product.Price, product.Quantity, product.Category)
	if err != nil {
		return err
	}
	fmt.Print(result.RowsAffected())
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = id
	return nil
}

// GetProduct retrieves a product by ID
func (ps *ProductStore) GetProduct(id int64) (*Product, error) {
	// TODO: Query the database for a product with the given ID
	// TODO: Return a Product struct populated with the data or an error if not found
	row := ps.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	product := Product{}
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category); err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func (ps *ProductStore) UpdateProduct(product *Product) error {
	// TODO: Update the product in the database
	// TODO: Return an error if the product doesn't exist
	_, err := ps.GetProduct(product.ID)
	if err != nil {
		return err
	}
	_, err = ps.db.Exec("UPDATE products SET name = ?, price = ?, quantity = ?, category = ? WHERE id = ?",
		product.Name,
		product.Price,
		product.Quantity,
		product.Category,
		product.ID)

	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct removes a product by ID
func (ps *ProductStore) DeleteProduct(id int64) error {
	// TODO: Delete the product from the database
	// TODO: Return an error if the product doesn't exist
	_, err := ps.GetProduct(id)
	if err != nil {
		return nil
	}
	_, err = ps.db.Exec("DELETE FROM products where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// ListProducts returns all products with optional filtering by category
func (ps *ProductStore) ListProducts(category string) ([]*Product, error) {
	// TODO: Query the database for products
	// TODO: If category is not empty, filter by category
	// TODO: Return a slice of Product pointers
	var rows *sql.Rows
	var err error
	if category == "" {
		rows, err = ps.db.Query("SELECT * from products")
	} else {
		rows, err = ps.db.Query("SELECT * from products WHERE category = ?", category)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*Product
	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

// BatchUpdateInventory updates the quantity of multiple products in a single transaction
func (ps *ProductStore) BatchUpdateInventory(updates map[int64]int) error {
	// TODO: Start a transaction
	// TODO: For each product ID in the updates map, update its quantity
	// TODO: If any update fails, roll back the transaction
	// TODO: Otherwise, commit the transaction
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for id, quantity := range updates {
		if _, err = ps.GetProduct(id); err != nil {
			tx.Rollback()
			return err
		}
		_, err := tx.Exec("UPDATE products SET quantity = ? WHERE id = ?", quantity, id)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
