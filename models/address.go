// model.go

package models

import (
	"database/sql"
)

// Address structure
type Address struct {
	ID         int    `json:"id"`
	Attention  string `json:"attention"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
}

// Save method
func (a *Address) Save(db *sql.DB, start, count int) (interface{}, error) {
	if a == nil {
		return nil, nil
	}

	// insert
	if a.ID == 0 {
		if a.IsEmpty() {
			return 0, nil
		}
		err := a.create(db)
		return a.ID, err
	}
	// delete
	if a.IsEmpty() {
		err := a.Delete(db)
		return nil, err
	}
	// update
	err := a.update(db)
	return a.ID, err
}

// // Read method
// func (a *Address) Read(db *sql.DB) error {
// 	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
// 		a.ID).Scan(&a.Name, &a.Price)
// }

// create method
func (a *Address) create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO addresses(attention, addr_line1, addr_line2, "+
		"city, addr_state, country, postal_code, phone, fax) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		a.Attention, a.Line1, a.Line2, a.City, a.State, a.Country, a.PostalCode, a.Phone, a.Fax)
	err = db.QueryRow("SELECT last_insert_rowid()").Scan(&a.ID)

	if err != nil {
		return err
	}

	return nil
}

// update method
func (a *Address) update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE addresses SET attention=$1, addr_line1=$2, addr_line2=$3, "+
			"city=$4, addr_state=$5, country=$6, postal_code=$7, phone=$8, fax=$9 "+
			"WHERE id=$10",
			a.Attention, a.Line1, a.Line2, a.City, a.State, a.Country, a.PostalCode, a.Phone, a.Fax, a.ID)

	return err
}

// Delete method
func (a *Address) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM addresses WHERE id=$1", a.ID)

	return err
}

// IsEmpty method
func (a *Address) IsEmpty() bool {
	return a.Attention == "" && a.Line1 == "" && a.Line2 == "" && a.City == "" && a.State == "" && a.Country == "" &&
		a.PostalCode == "" && a.Phone == "" && a.Fax == ""
}
