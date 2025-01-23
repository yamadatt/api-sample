package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"yamadatt/api-sample/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/stocks", handlers.RegisterStockHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
