# このリポジトリは？

golangのAPIサンプル。以下のAPIを実装。

- 在庫を登録
- 在庫をチェック
- 売上を上げる
- 売上をチェック
- データを全て削除

golangで実装しており、ポート8080で動作する。

使い方のサンプルコマンドは[使用例](#使用例)参照。

- [このリポジトリは？](#このリポジトリは)
  - [API Documentation](#api-documentation)
    - [Register Stock](#register-stock)
      - [Endpoint](#endpoint)
      - [Request Parameters](#request-parameters)
      - [Example Request](#example-request)
      - [Example Response](#example-response)
    - [Check Stock](#check-stock)
      - [Endpoint](#endpoint-1)
      - [Request Parameters](#request-parameters-1)
      - [Example Request (with name)](#example-request-with-name)
      - [Example Response (with name)](#example-response-with-name)
      - [Example Request (without name)](#example-request-without-name)
      - [Example Response (without name)](#example-response-without-name)
    - [Register Sales](#register-sales)
      - [Endpoint](#endpoint-2)
      - [Request Parameters](#request-parameters-2)
      - [Example Request](#example-request-1)
      - [Example Response](#example-response-1)
    - [Truncate Stocks and Sales Tables](#truncate-stocks-and-sales-tables)
      - [Endpoint](#endpoint-3)
      - [Example Response](#example-response-2)
  - [使用例](#使用例)
    - [登録](#登録)
    - [在庫チェック](#在庫チェック)
    - [売り上げる](#売り上げる)
      - [単価なし](#単価なし)
      - [単価あり](#単価あり)
    - [売上チェック](#売上チェック)
    - [削除](#削除)



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

## 使用例

### 登録

```bash
curl -v -d '{"name": "product1","amount": 100}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/stocks
```

### 在庫チェック

```bash
curl http://192.168.1.78:8080/v1/stocks/product1
```
### 売り上げる

#### 単価なし

```bash
curl -v -d '{"name": "product1","amount": 4}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/sales
```
#### 単価あり

```bash
curl -v -d '{"name": "product1","amount": 4,"price": 2000}' -H 'Content-Type: application/json' http://192.168.1.78:8080/v1/sales
```
### 売上チェック

```bash
curl http://192.168.1.78:8080/v1/sales
```

### 削除

```bash
curl -X DELETE http://192.168.1.78:8080/v1/stocks
```