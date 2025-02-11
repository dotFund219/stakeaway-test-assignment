# Use official Golang image
FROM golang:1.23.4 AS builder

WORKDIR /app

# Install SQLite3 dependencies
RUN apt-get update && apt-get install -y gcc libc6-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable CGO for SQLite3
ENV CGO_ENABLED=1

RUN go build -o main ./cmd/main.go

# Final lightweight image
FROM debian:latest

WORKDIR /root/

# Copy built binary from the builder stage
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Run the binary
CMD ["./main"]
