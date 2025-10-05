# Cart REST Service

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

# Cart Microservice

Shopping cart management microservice for The Cheeky Cart online store. Handles adding products to cart, updating quantities, and managing cart state before checkout.

## Project Structure

```
cart-rest/
├── main.go
├── internal/
│   ├── cart.go         # Cart model and methods
│   ├── service.go      # Business logic
│   ├── handler.go      # HTTP handlers
│   ├── telemetry.go    # OpenTelemetry setup
│   └── middleware.go   # Metrics middleware
├── k6-tests/           # Performance tests
│   ├── smoke-test.js
│   ├── load-test.js
│   ├── stress-test.js
│   ├── spike-test.js
│   └── run-all-tests.sh
├── go.mod
└── go.sum
```

## API Endpoints

### Cart Operations

- **GET /cart/:id** - Get customer's cart
- **POST /cart/:id/items** - Add product to cart
- **PUT /cart/:id/items/:productId** - Update product quantity
- **DELETE /cart/:id/items/:productId** - Remove product from cart

### Monitoring

- **GET /health** - Health check endpoint
- **GET /metrics** - Prometheus metrics

## Usage

### Run

```bash
make run
```

Server starts on `http://localhost:8080`

### Build

```bash
make build
```

Binary created in `bin/cart-rest`

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

## Observability

### Metrics

The service exposes Prometheus metrics at `/metrics`:

- `http_requests_total` - Total HTTP requests (counter)
- `http_request_duration_seconds` - Request latency (histogram)

All metrics include labels: `method`, `path`, `status_code`

### View Metrics

```bash
curl http://localhost:8080/metrics
```

## Performance Testing

The service includes k6 performance tests in `k6-tests/` directory.

### Run Tests

```bash
# Individual tests
make test-smoke    # 4 min - Basic validation
make test-load     # 16 min - Normal load
make test-stress   # 22 min - Breaking point
make test-spike    # 5 min - Traffic spikes

# All tests
make test-all      # ~47 min
```
