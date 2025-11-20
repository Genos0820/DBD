# Coditas API

This is a REST API built using the [Gin Web Framework] in Go. The API includes custom validation for PAN and mobile numbers and provides an endpoint to create a user.

## Features

- **Custom Validators**:
  - PAN validation: Ensures the PAN is in the correct format (e.g., `ABCDE1234F`).
  - Mobile validation: Ensures the mobile number is exactly 10 digits.
- **Strict Validation**:
  - Disallows extra fields in the request payload.
- **Middleware**:
  - Latency logger middleware to log request latency.
- **Unit Tests**:
  - Comprehensive test cases for the `/create-user` endpoint.

## Prerequisites

- [Go](https://golang.org/) (version 1.18 or higher)
- A terminal or IDE to run the application

## Install Dependencies
go mod tidy

## Usage
go run main.go

curl -X POST http://localhost:8080/create-user \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "pan": "ABCDE1234F",
    "mobile": "9876543210",
    "email": "john.doe@example.com"
}'

# PowerShell equivalent
Invoke-WebRequest -Uri "http://localhost:8080/create-user" `
  -Method POST `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{
    "name": "John Doe",
    "pan": "ABCDE1234F",
    "mobile": "9876543210",
    "email": "john.doe@example.com"
  }'

## API Endpoints
POST /create-user
-Description: Creates a new user.
-Request Body:
{
    "name": "string (required)",
    "pan": "string (required, custom validation)",
    "mobile": "string (required, custom validation)",
    "email": "string (required, valid email)"
}
-Response:
    -Success: 200 OK
    {
        "message": "request payload completed successfully!"
    }
    -Error: 400 Bad Request
    {
        "error_message": "Validation or parsing error message"
    }

## Project Structure
CoditasApi/
├── services/
│   └── service.go       # Contains the handler logic for the API
├── middlewares/
│   └── logger.go        # Middleware for logging request latency
├── main.go              # Main entry point for the application
├── main_test.go         # Unit tests for the server
├── go.mod               # Go module file
└── go.sum               # Dependency lock file