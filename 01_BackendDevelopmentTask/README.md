# Golang API Backend with Gin Framework

## Overview

This project is a Golang API backend built using the Gin framework, featuring an in-memory database and Docker support.

## Project Structure

```
├── cmd/ # Entry point for the application
│ ├── main.go # Main entry point
├── config/ # Configuration files
│ ├── config.yaml # App configuration
├── internal/ # Internal application logic
│ ├── handlers/ # HTTP handlers
│ ├── models/ # Data models
│ ├── services/ # Business logic
│ ├── repositories/ # Data access layer
│ ├── middlewares/ # Middleware functions
│ ├── utils/ # Utility functions
├── pkg/ # Shared packages
├── db/ # In-memory database setup
├── tests/ # Unit and integration tests
├── Dockerfile # Docker configuration
├── docker-compose.yml # Docker Compose setup
├── go.mod # Go module file
├── go.sum # Dependencies lock file
├── Makefile # Makefile for automation
├── README.md # Documentation
```

## Prerequisites

Make sure you have the following installed:

- Go (v1.21 or later)

- Docker & Docker Compose

## Installation & Setup

### 1. Clone the Repository

```
git clone https://github.com/yourusername/golang-api-backend.git
cd golang-api-backend
```

### 2. Run with Docker

```
docker-compose up --build
```

### 3. Run Locally (Without Docker)

If you want to run the app locally without Docker:

```
go mod tidy  # Install dependencies
go run cmd/main.go
```

## API Endpoints

| Method | Endpoint | Description                                                                         |
| ------ | -------- | ----------------------------------------------------------------------------------- |
| POST   | /stake   | Simulates staking by logging the transaction in a local database                    |
| GET    | /reward  | Simulate rewards by calculating 5% of the staked amount.                            |
| GET    | /health  | Returns the status of the service (e.g., database connectivity and service uptime). |

## Testing the API

### Stake Test

```
curl --location 'http://localhost:8080/stake' \
--header 'Content-Type: application/json' \
--data '{
    "wallet_address": "0x1ed5A84F44F88b00eEA922d24CA6db8Ad04436cb",
    "amount": 4
}'
```

Expected Output

```
{
    "message": "Staking successful"
}
```

### Rewards Test

```
curl --location 'http://localhost:8080/rewards/0x1ed5A84F44F88b00eEA922d24CA6db8Ad04436cb'
```

Expected Output

```
{
    "rewards": 0.2,
    "wallet_address": "0x1ed5A84F44F88b00eEA922d24CA6db8Ad04436cb"
}
```

### Health Test

```
curl --location 'http://localhost:8080/health'
```

Expected Output

```
{
    "status": "Service is healthy",
    "uptime": "4m37.306106072s"
}
```

## Stopping the Server

To stop the Docker container:

```
docker-compose down
```

## License

This project is licensed under the MIT License.
