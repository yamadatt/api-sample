package handlers

import (
	"api-sample/database"
	"api-sample/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func RegisterStockHandler(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if stock.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	// Validate name field
	isValidName := regexp.MustCompile(`^[a-z]{1,8}$`).MatchString
	if !isValidName(stock.Name) {
		http.Error(w, `{"message": "ERROR"}`, http.StatusBadRequest)
		return
	}

	if stock.Amount == 0 {
		stock.Amount = 1
	}

	db, err := database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = stock.Register()
	if err != nil {
		fmt.Println("debug:Failed to register stock:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Stock registered successfully",
		"product": stock,
	})
}

func GetStockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	db, err := database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if name != "" {
		stock, err := database.GetStockByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(stock)
	} else {
		stocks, err := database.GetAllStocks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(stocks)
	}
}

func DeleteStocksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.TruncateTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Stocks and sales tables truncated successfully",
	})
}
