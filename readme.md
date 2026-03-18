# Routing Number API Service

A lightweight Go microservice that provides bank information lookup by routing number. This API returns detailed information about U.S. financial institutions based on their ABA routing numbers.

## Features

- Fast routing number lookup
- Comprehensive bank information including address and contact details
- Phone number formatting
- CORS support for web applications
- Health check endpoint for load balancers
- Lightweight Docker container
- Production-ready deployment support

## Prerequisites

- Go 1.18 or higher
- Docker (for containerized deployment)

## Local Development

### Running Locally

```bash
# Install dependencies
go mod download

# Run the service
go run main.go

# Or specify a custom data file
go run main.go path/to/banks.json
```

The service will start on `http://localhost:8080`

### Building

```bash
go build -o routing-number-api
./routing-number-api
```

### Generating Swagger Documentation

If you make changes to the API annotations, regenerate the Swagger docs:

```bash
swag init
```

This will update the `docs/` directory with the latest API documentation.

## API Documentation

### Interactive Swagger Documentation

Once the service is running, you can access the interactive Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

The Swagger UI provides:
- Interactive API testing
- Complete API schema documentation
- Request/response examples
- Easy-to-use interface for exploring all endpoints

### Endpoints

### GET /banks

Returns a list of all banks in JSON format.

**Example Response:**
```json
[
  {
    "routing_number": "324184440",
    "bank": "IDAHO UNITED CREDIT UNION",
    "address": "PO BOX 2268",
    "city": "BOISE",
    "state": "ID",
    "zip": "83701",
    "zip4": "2268",
    "phone": "2083882138"
  },
  {
    "routing_number": "324274033",
    "bank": "ELKO FEDERAL CREDIT UNION",
    "address": "2397 MTN CITY HWY",
    "city": "EIKO",
    "state": "NV",
    "zip": "89801",
    "zip4": "0000",
    "phone": "7757384083"
  }
]
```

### GET /banks/:routing

Returns bank information for the specified routing number.

**Example Request:**
```
GET /banks/031101334
```

**Example Response:**
```json
{
  "routing_number": "031101334",
  "bank": "SoFi Bank, N.A.",
  "address": "San Francisco, CA",
  "city": "San Francisco",
  "state": "CA",
  "zip": "",
  "phone": "(855) 936-2269",
  "message": "OK"
}
```

**Error Response (404):**
```json
{
  "message": "bank not found"
}
```

### GET /health

Health check endpoint for load balancers and monitoring.

**Example Response:**
```json
"up"
```

## Docker Deployment

### Building the Docker Image

```bash
docker build -t routing-number-api .
```

### Running with Docker

```bash
docker run -p 8080:8080 routing-number-api
```


## Updating Bank Data

To add or update routing number information:

1. Edit the `data/banks.json` file with the new bank routing information
2. Rebuild and redeploy using the deployment method of your choice

### Data Format

```json
{
  "routing_number": "123456789",
  "bank": "BANK NAME",
  "address": "STREET ADDRESS",
  "city": "CITY",
  "state": "ST",
  "zip": "12345",
  "zip4": "1234",
  "phone": "1234567890"
}
```

## Technology Stack

- **Language:** Go 1.18+
- **Web Framework:** Gin
- **Phone Formatting:** libphonenumber
- **API Documentation:** Swagger/OpenAPI
- **Containerization:** Docker (Alpine-based multi-stage build)
