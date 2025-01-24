package tests

import (
	"api-sample/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegisterStockHandler(t *testing.T) {
	reqBody := []byte(`{"name": "Product A", "amount": 10}`)
	req, err := http.NewRequest("POST", "/v1/stocks", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.RegisterStockHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Stock registered successfully", response["message"])
	assert.Equal(t, "Product A", response["product"].(map[string]interface{})["name"])
	assert.Equal(t, float64(10), response["product"].(map[string]interface{})["amount"])
}

func TestGetStockHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/stocks/Product A", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/stocks/{name}", handlers.GetStockHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var stock map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&stock)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Product A", stock["name"])
	assert.Equal(t, float64(10), stock["amount"])
}

func TestRegisterSalesHandler(t *testing.T) {
	reqBody := []byte(`{"name": "Product A", "amount": 2, "price": 50.0}`)
	req, err := http.NewRequest("POST", "/v1/sales", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.RegisterSalesHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Sale registered successfully", response["message"])
	assert.Equal(t, "Product A", response["product"])
	assert.Equal(t, float64(2), response["amount"])
	assert.Equal(t, float64(100), response["totalPrice"])
}

func TestGetSalesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/sales", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetSalesHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var sales []map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&sales)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, sales)
	assert.Equal(t, float64(100), sales[0]["total_price"])
}

func TestDeleteStocksHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/stocks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteStocksHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Stocks and sales tables truncated successfully", response["message"])
}
