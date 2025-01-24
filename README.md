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



## example

登録

curl -v -d '{"name": "product1","amount": 100}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/stocks

在庫チェック

curl http://192.168.1.78:8080/v1/stocks/product1