package handlers

import (
	"api-sample/database"
	"api-sample/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	if stock.Amount == 0 {
		stock.Amount = 1
	}

	fmt.Println("debug:Registering stock:", stock)

	//dbのオープンとクローズをここで行う
	//dbのオープンとクローズをここで行う
	//dbのオープンとクローズをここで行う
	err = database.InitDB("mydb.db")
	if err != nil {
		log.Fatal(err)
	}

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
