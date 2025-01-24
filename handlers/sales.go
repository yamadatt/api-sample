package handlers

import (
	"api-sample/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	if sale.Amount == 0 {
		sale.Amount = 1
	}

	fmt.Println("debug:Registering sale:", sale)

	err = database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}

	stock, err := database.GetStockByName(sale.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if stock.Amount < sale.Amount {
		http.Error(w, "insufficient stock", http.StatusBadRequest)
		return
	}

	err = database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	stock.Amount -= sale.Amount
	err = stock.Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPrice := 0.0
	if sale.Price > 0 {
		totalPrice = float64(sale.Amount) * sale.Price
	}

	query := `INSERT INTO sales (name, amount, total_price) VALUES (?, ?, ?)`
	_, err = db.Exec(query, sale.Name, sale.Amount, totalPrice)
	if err != nil {
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
