package tests

import (
	"api-sample/models"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	models.InitDB(":memory:")

	stock := models.Stock{Name: "product1", Amount: 10}
	err := stock.Register()
	assert.NoError(t, err)

	registeredStock, err := models.GetStockByName("product1")
	assert.NoError(t, err)
	assert.Equal(t, "product1", registeredStock.Name)
	assert.Equal(t, 10, registeredStock.Amount)
}
