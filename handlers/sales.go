package handlers

import (
	"api-sample/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func RegisterSalesHandler(w http.ResponseWriter, r *http.Request) {
	var sale struct {
		Name   string  `json:"name"`
		Amount int     `json:"amount"`
		Price  float64 `json:"price"`
	}

	err := json.NewDecoder(r.Body).Decode(&sale)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if sale.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	// Validate name field
	isValidName := regexp.MustCompile(`^[a-z]{1,8}$`).MatchString
	if !isValidName(sale.Name) {
		http.Error(w, `{"message": "ERROR"}`, http.StatusBadRequest)
		return
	}

	fmt.Println("debug:Registering sale:", sale.Name)

	if sale.Amount == 0 {
		sale.Amount = 1
	}

	fmt.Println("debug:Registering sale:", sale)

	db, err := database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stock, err := database.GetStockByName(sale.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if stock.Amount < sale.Amount {
		http.Error(w, "stockが不足しているよ！！", http.StatusBadRequest)
		return
	}

	stock.Amount -= sale.Amount
	err = database.UpdateStock(stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPrice := 0.0
	if sale.Price > 0 {
		totalPrice = float64(sale.Amount) * sale.Price
	}

	var existingSale struct {
		TotalPrice float64 `json:"total_price"`
	}

	query := `SELECT total_price FROM sales WHERE name = ?`
	err = db.QueryRow(query, sale.Name).Scan(&existingSale.TotalPrice)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPrice += existingSale.TotalPrice

	query = `INSERT INTO sales (name, amount, total_price) VALUES (?, ?, ?)
			 ON CONFLICT(name) DO UPDATE SET amount = amount + excluded.amount, total_price = total_price + excluded.total_price`
	_, err = db.Exec(query, sale.Name, sale.Amount, totalPrice)

	fmt.Println("debug:Sale registered successfully??")
	if err != nil {
		fmt.Print("debug:Failed to register sale:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Sale registered successfully",
		"product":    sale.Name,
		"amount":     sale.Amount,
		"totalPrice": totalPrice,
	})
}

func GetSalesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sales, err := database.GetAllSales()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}
