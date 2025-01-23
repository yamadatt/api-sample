package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		amount INTEGER NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}

func InsertStock(name string, amount int) error {
	query := `INSERT INTO stocks (name, amount) VALUES (?, ?)`
	_, err := db.Exec(query, name, amount)
	if err != nil {
		return fmt.Errorf("failed to insert stock: %v", err)
	}

	return nil
}
