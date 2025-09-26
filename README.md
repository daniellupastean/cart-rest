# Cart Microservice

Shopping cart management microservice for The Cheeky Cart online store. Handles adding products to cart, updating quantities, and managing cart state before checkout.

## Project Structure

```
cart-rest/
├── main.go
├── internal/
│   ├── cart.go     # Cart model and methods
│   ├── service.go  # Business logic
│   └── handler.go  # HTTP handlers
├── go.mod
└── go.sum
```

## API Endpoints

- **GET /cart/:id** - Get customer's cart
- **POST /cart/:id/items** - Add product to cart
- **PUT /cart/:id/items/:productId** - Update product quantity
- **DELETE /cart/:id/items/:productId** - Remove product from cart

## Usage

### Run

```bash
make run
```

Server starts on `http://localhost:8080`

## Examples

### Add product to cart

```bash
curl -X POST http://localhost:8080/cart/user-001/items \
  -H "Content-Type: application/json" \
  -d '{"productId": "prod-001", "quantity": 2}'
```

### Get customer cart

```bash
curl -X GET http://localhost:8080/cart/user-001
```

### Update product quantity

```bash
curl -X PUT http://localhost:8080/cart/user-001/items/prod-001 \
  -H "Content-Type: application/json" \
  -d '{"quantity": 5}'
```

### Remove product from cart

```bash
curl -X DELETE http://localhost:8080/cart/user-001/items/prod-001
```
