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
	db, err := sql.Open("sqlite3", "./stock.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	query := `INSERT INTO stocks (name, amount) VALUES (?, ?)`
	_, err = db.Exec(query, s.Name, s.Amount)
	if err != nil {
		return fmt.Errorf("failed to register stock: %v", err)
	}

	return nil
}
