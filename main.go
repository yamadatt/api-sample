package main

import (
	"api-sample/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/stocks", handlers.RegisterStockHandler).Methods("POST")
	r.HandleFunc("/v1/stocks", handlers.GetStockHandler).Methods("GET")
	r.HandleFunc("/v1/stocks/{name}", handlers.GetStockHandler).Methods("GET")
	r.HandleFunc("/v1/sales", handlers.RegisterSalesHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
