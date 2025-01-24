package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Stock struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func (s *Stock) Register() error {
	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("debug:Connected to database")

	query := `INSERT INTO stocks (name, amount) VALUES (?, ?)
			  ON CONFLICT(name) DO UPDATE SET amount = amount + excluded.amount`
	_, err = db.Exec(query, s.Name, s.Amount)
	if err != nil {
		return fmt.Errorf("failed to register stock: %v", err)
	}

	return nil
}
