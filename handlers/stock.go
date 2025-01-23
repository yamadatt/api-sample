package handlers

import (
	"encoding/json"
	"net/http"
	"yamadatt/api-sample/models"
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

	err = stock.Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Stock registered successfully",
		"product": stock,
	})
}
