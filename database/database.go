package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Stock represents a stock item in the database
type Stock struct {
	Name   string
	Amount int
}

var db *sql.DB

func InitDB(dataSourceName string) error {
	fmt.Println("debug:InitDB", dataSourceName)

	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		amount INTEGER NOT NULL
	)`

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("debug:stocks exist", dataSourceName)

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

func GetStockByName(name string) (*Stock, error) {
	query := `SELECT name, amount FROM stocks WHERE name = ?`
	row := db.QueryRow(query, name)

	var stock Stock
	err := row.Scan(&stock.Name, &stock.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no stock found for product: %s", name)
		}
		return nil, fmt.Errorf("failed to retrieve stock: %v", err)
	}

	return &stock, nil
}

func GetAllStocks() ([]Stock, error) {
	query := `SELECT name, amount FROM stocks ORDER BY name ASC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve stocks: %v", err)
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.Name, &stock.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stock: %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
