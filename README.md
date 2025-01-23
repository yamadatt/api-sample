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
