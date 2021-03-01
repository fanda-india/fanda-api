// model.go

package models

import (
	"database/sql"
)

// Product structure
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// List method
func (p *Product) List(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM products ORDER BY id LIMIT $1 OFFSET $2",
		count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// Read method
func (p *Product) Read(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

// Create method
func (p *Product) Create(db *sql.DB) error {
	// "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",

	// err := db.QueryRow(
	// 	"INSERT INTO products(name, price) VALUES($1, $2);SELECT last_insert_rowid()",
	// 	p.Name, p.Price).Scan(&p.ID)

	_, err := db.Exec("INSERT INTO products(name, price) VALUES($1, $2)", p.Name, p.Price) //.Scan(&p.ID)
	err = db.QueryRow("SELECT last_insert_rowid()").Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

// Update method
func (p *Product) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

// Delete method
func (p *Product) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}
