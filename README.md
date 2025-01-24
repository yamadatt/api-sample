# api-sample
golangのAPIサンプル

## API Documentation

### Register Stock

#### Endpoint
- **HTTP Method:** POST
- **Request URI:** /v1/stocks

#### Request Parameters
- **name (required):** The name of the product
- **amount (optional):** The quantity of the product to add to the stock (positive integer). Default is 1.

#### Example Request
```json
{
  "name": "Product A",
  "amount": 10
}
```

#### Example Response
- **HTTP Status Code:** 200 OK
- **Response Body:**
```json
{
  "message": "Stock registered successfully",
  "product": {
    "name": "Product A",
    "amount": 10
  }
}
```

### Check Stock

#### Endpoint
- **HTTP Method:** GET
- **Request URI:** /v1/stocks(/:name)　（:nameは対象の商品の名前）

#### Request Parameters
- **name (optional):** The name of the product to check the stock for. If not specified, returns all products sorted by name in ascending order.

#### Example Request (with name)
- **Request URI:** /v1/stocks/ProductA

#### Example Response (with name)
- **HTTP Status Code:** 200 OK
- **Response Body:**
```json
{
  "name": "Product A",
  "amount": 10
}
```

#### Example Request (without name)
- **Request URI:** /v1/stocks

#### Example Response (without name)
- **HTTP Status Code:** 200 OK
- **Response Body:**
```json
[
  {
    "name": "Product A",
    "amount": 10
  },
  {
    "name": "Product B",
    "amount": 5
  }
]
```

### Register Sales

#### Endpoint
- **HTTP Method:** POST
- **Request URI:** /v1/sales

#### Request Parameters
- **name (required):** The name of the product
- **amount (optional):** The quantity of the product to sell (positive integer). Default is 1.
- **price (optional):** The price of the product (positive number). If provided, the total price (price * amount) will be added to the sales total.

#### Example Request
```json
{
  "name": "Product A",
  "amount": 2,
  "price": 50.0
}
```

#### Example Response
- **HTTP Status Code:** 200 OK
- **Response Body:**
```json
{
  "message": "Sale registered successfully",
  "product": "Product A",
  "amount": 2,
  "totalPrice": 100.0
}
```

### Truncate Stocks and Sales Tables

#### Endpoint
- **HTTP Method:** DELETE
- **Request URI:** /v1/stocks

#### Example Response
- **HTTP Status Code:** 200 OK
- **Response Body:**
```json
{
  "message": "Stocks and sales tables truncated successfully"
}
```

## example

登録

curl -v -d '{"name": "product1","amount": 100}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/stocks

在庫チェック

curl http://192.168.1.78:8080/v1/stocks/product1

売り上げる

単価なし

curl -v -d '{"name": "product1","amount": 4}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/sales

単価あり

curl -v -d '{"name": "product1","amount": 4,"price": 2000}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/sales

売上チェック

curl http://192.168.1.78:8080/v1/sales
