package tests

import (
	"api-sample/database"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	db, err := database.InitDB(":memory:")
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestInsertStock(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	err := database.InsertStock("product1", 10)
	assert.NoError(t, err)

	stock, err := database.GetStockByName("product1")
	assert.NoError(t, err)
	assert.Equal(t, "product1", stock.Name)
	assert.Equal(t, 10, stock.Amount)
}

func TestGetStockByName(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	database.InsertStock("product1", 10)

	stock, err := database.GetStockByName("product1")
	assert.NoError(t, err)
	assert.Equal(t, "product1", stock.Name)
	assert.Equal(t, 10, stock.Amount)
}

func TestGetAllStocks(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	database.InsertStock("product1", 10)
	database.InsertStock("product2", 20)

	stocks, err := database.GetAllStocks()
	assert.NoError(t, err)
	assert.Len(t, stocks, 2)
}

func TestUpdateStock(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	database.InsertStock("product1", 10)

	stock, _ := database.GetStockByName("product1")
	stock.Amount = 15
	err := database.UpdateStock(stock)
	assert.NoError(t, err)

	updatedStock, err := database.GetStockByName("product1")
	assert.NoError(t, err)
	assert.Equal(t, 15, updatedStock.Amount)
}

func TestGetAllSales(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	query := `INSERT INTO sales (name, amount, total_price) VALUES (?, ?, ?)`
	db.Exec(query, "product1", 1, 100.0)

	sales, err := database.GetAllSales()
	assert.NoError(t, err)
	assert.Len(t, sales, 1)
	assert.Equal(t, 100.0, sales[0]["total_price"])
}

func TestTruncateTables(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	defer db.Close()

	database.InsertStock("product1", 10)
	query := `INSERT INTO sales (name, amount, total_price) VALUES (?, ?, ?)`
	db.Exec(query, "product1", 1, 100.0)

	err := database.TruncateTables()
	assert.NoError(t, err)

	stocks, _ := database.GetAllStocks()
	assert.Len(t, stocks, 0)

	sales, _ := database.GetAllSales()
	assert.Len(t, sales, 0)
}
