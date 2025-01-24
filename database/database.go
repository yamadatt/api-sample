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

func InitDB(dataSourceName string) (*sql.DB, error) {
	fmt.Println("debug:InitDB", dataSourceName)

	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		amount INTEGER NOT NULL
	)`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	query = `
	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		amount INTEGER NOT NULL,
		total_price REAL NOT NULL
	)`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("debug:stocks and sales exist", dataSourceName)

	return db, nil
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

func UpdateStock(stock *Stock) error {
	query := `UPDATE stocks SET amount = ? WHERE name = ?`
	_, err := db.Exec(query, stock.Amount, stock.Name)
	if err != nil {
		return fmt.Errorf("failed to update stock: %v", err)
	}

	return nil
}

func GetAllSales() ([]map[string]interface{}, error) {
	// query := `SELECT name, amount, ROUND(total_price, 2) as total_price FROM sales`
	query := `SELECT SUM(total_price) as total_sales FROM sales`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sales: %v", err)
	}
	defer rows.Close()

	var sales []map[string]interface{}
	for rows.Next() {
		// var name string
		// var amount int
		var totalPrice float64
		// err := rows.Scan(&name, &amount, &totalPrice)
		err := rows.Scan(&totalPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to scan sale: %v", err)
		}
		sale := map[string]interface{}{
			// "name":        name,
			// "amount":      amount,
			"total_price": totalPrice,
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
